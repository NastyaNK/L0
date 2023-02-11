package psql

import (
	"L0/internal/entity"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func NewDB(connect *entity.DBConnect) (*Postgres, error) {
	psql := Postgres{}
	_, err := psql.Connect(connect.User, connect.Password, connect.DBName, connect.Host)
	return &psql, err
}
func (psql *Postgres) Connect(user, password, base, host string) (*sqlx.DB, error) {
	var err error
	if psql.db == nil {
		psql.db, err = sqlx.Connect("postgres",
			fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, base, host))
	}
	return psql.db, err
}
func (psql *Postgres) GetAll() ([]*entity.Model, error) {
	var models []*entity.Model
	rows, err := psql.db.Queryx("SELECT order_uid FROM orders")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var str string
	for rows.Next() {
		err = rows.Scan(&str)
		if err != nil {
			log.Println(err)
			return models, err
		}
		model, err := psql.Get(str)
		if err != nil {
			return models, err
		}
		models = append(models, model)
	}
	return models, nil
}

func (psql *Postgres) Get(id string) (*entity.Model, error) {
	var order entity.Order
	err := psql.db.Get(&order, "select * from orders where order_uid=$1", id)
	if err != nil {
		log.Println("ORDER DB:", err)
		return nil, err
	}
	var items []entity.Item
	err = psql.db.Select(&items, "SELECT * FROM items where order_uid=$1", id)
	if err != nil {
		log.Println("ITEMS DB:", err)
		return nil, err
	}
	var delivery entity.Delivery
	err = psql.db.Get(&delivery, "select * from delivery where phone=$1", order.Delivery)
	if err != nil {
		log.Println("DELIVERY DB:", err)
		return nil, err
	}
	var payment entity.Payment
	err = psql.db.Get(&payment, "select * from payment where transaction=$1", order.Payment)
	if err != nil {
		log.Println("PAYMENT DB:", err)
		return nil, err
	}
	return order.ToModel(&payment, &delivery, &items), nil
}
func (psql *Postgres) Set(model *entity.Model) error {
	order, payment, delivery, items := model.ToDB()
	tx := psql.db.MustBegin()
	_, err := tx.NamedExec("insert into delivery (phone, name, zip, city, address, region, email) values "+
		"(:phone, :name, :zip, :city, :address, :region, :email);", &delivery)
	if err != nil {
		log.Println("DELIVERY SET DB:", err)
		return err
	}
	_, err = tx.NamedExec("insert into payment (transaction, request_id, currency, provider, amount, "+
		"payment_dt, bank, delivery_cost,goods_total, custom_fee) values (:transaction, :request_id, :currency, "+
		":provider, :amount, :payment_dt, :bank, :delivery_cost,:goods_total, :custom_fee);", &payment)
	if err != nil {
		log.Println("PAYMENT SET DB:", err)
		return err
	}
	_, err = tx.NamedExec("insert into orders (order_uid, track_number, entry, delivery, payment, locale, "+
		"internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) values "+
		"(:order_uid, :track_number, :entry, :delivery, :payment, :locale, :internal_signature, :customer_id, "+
		":delivery_service, :shardkey, :sm_id, :date_created, :oof_shard);", &order)
	if err != nil {
		log.Println("ORDERS SET DB:", err)
		return err
	}

	for _, item := range *items {
		_, err := tx.NamedExec("insert into items (chrt_id, track_number, price, rid, name, sale, size, "+
			"total_price, nm_id, brand, status, order_uid) values (:chrt_id, :track_number, :price, :rid, :name, "+
			":sale, :size, :total_price, :nm_id, :brand, :status, :order_uid);", &item)
		if err != nil {
			log.Println("ITEMS SET DB:", err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println("COMMIT SET DB:", err)
	}
	return err
}

func GetDBConfig(file string) (*entity.DBConnect, error) {
	var db entity.DBConnect
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &db)
	return &db, err
}
