package psql

import (
	"L0/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Postgres struct {
	db *sqlx.DB
}

func NewDB(user, password, base, host string) *Postgres {
	psql := Postgres{}
	_, err := psql.Connect(user, password, base, host)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &psql
}
func (psql *Postgres) Connect(user, password, base, host string) (*sqlx.DB, error) {
	var err error
	if psql.db == nil {
		psql.db, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, base, host))
	}
	return psql.db, err
}
func (psql *Postgres) GetAll() []*entity.Model {
	var models []*entity.Model
	rows, err := psql.db.Queryx("SELECT order_uid FROM orders")
	var str string
	for rows.Next() {
		err = rows.Scan(&str)
		if err != nil {
			log.Fatalln(err)
		}
		models = append(models, psql.Get(str))
	}
	return models
}

func (psql *Postgres) Get(id string) *entity.Model {
	var order entity.Order
	err := psql.db.Get(&order, "select * from orders where order_uid=$1", id)
	if err != nil {
		fmt.Println("ORDER DB:", err)
		return nil //вытаскиваем заказы из базы
	}
	var items []entity.Item
	err = psql.db.Select(&items, "SELECT * FROM items where order_uid=$1", id)
	if err != nil {
		fmt.Println("ITEMS DB:", err)
		return nil
	}
	var delivery entity.Delivery
	err = psql.db.Get(&delivery, "select * from delivery where phone=$1", order.Delivery)
	if err != nil {
		fmt.Println("DELIVERY DB:", err)
		return nil
	}
	var payment entity.Payment
	err = psql.db.Get(&payment, "select * from payment where transaction=$1", order.Payment)
	if err != nil {
		fmt.Println("PAYMENT DB:", err)
		return nil
	}
	return order.ToModel(&payment, &delivery, &items)
}
func (psql *Postgres) Set(model *entity.Model) {
	order, payment, delivery, items := model.ToDB()
	tx := psql.db.MustBegin()
	_, err := tx.NamedExec("insert into delivery (phone, name, zip, city, address, region, email) values (:phone, :name, :zip, :city, :address, :region, :email);", &delivery)
	if err != nil {
		log.Fatal("DELIVERY SET DB:", err)
		return
	}
	_, err = tx.NamedExec("insert into payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost,goods_total, custom_fee) values (:transaction, :request_id, :currency, :provider, :amount, :payment_dt, :bank, :delivery_cost,:goods_total, :custom_fee);", &payment)
	if err != nil {
		log.Fatal("PAYMENT SET DB:", err)
		return
	}
	_, err = tx.NamedExec("insert into orders (order_uid, track_number, entry, delivery, payment, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)values(:order_uid, :track_number, :entry, :delivery, :payment, :locale, :internal_signature, :customer_id, :delivery_service, :shardkey, :sm_id, :date_created, :oof_shard);", &order)
	if err != nil {
		log.Fatal("ORDERS SET DB:", err)
		return
	}

	for _, item := range *items {
		_, err := tx.NamedExec("insert into items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid) "+
			"values (:chrt_id, :track_number, :price, :rid, :name, :sale, :size, :total_price, :nm_id, :brand, :status, :order_uid);", &item)
		if err != nil {
			log.Fatal("ITEMS SET DB:", err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal("COMMIT SET DB:", err)
		return
	}

}
