package entity

import "time"

type Model struct {
	OrderUid          string        `json:"order_uid"`
	TrackNumber       string        `json:"track_number"`
	Entry             string        `json:"entry"`
	Delivery          ModelDelivery `json:"delivery"`
	Payment           ModelPayment  `json:"payment"`
	Items             []ModelItem   `json:"items"`
	Locale            string        `json:"locale"`
	InternalSignature string        `json:"internal_signature"`
	CustomerId        string        `json:"customer_id"`
	DeliveryService   string        `json:"delivery_service"`
	Shardkey          string        `json:"shardkey"`
	SmId              int           `json:"sm_id"`
	DateCreated       time.Time     `json:"date_created"`
	OofShard          string        `json:"oof_shard"`
}
type ModelItem struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}
type ModelPayment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}
type ModelDelivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (m *Model) ToDB() (*Order, *Payment, *Delivery, *[]Item) {
	order := Order{
		m.OrderUid,
		m.TrackNumber,
		m.Entry,
		m.Delivery.Phone,
		m.Payment.Transaction,
		m.Locale,
		m.InternalSignature,
		m.CustomerId,
		m.DeliveryService,
		m.Shardkey,
		m.SmId,
		m.DateCreated,
		m.OofShard,
	}

	payment := Payment(m.Payment)
	delivery := Delivery(m.Delivery)
	items := make([]Item, len(m.Items))
	for i, item := range m.Items {
		items[i] = Item{
			item.ChrtId,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmId,
			item.Brand,
			item.Status,
			m.OrderUid,
		}
	}
	return &order, &payment, &delivery, &items
}
func (order *Order) ToModel(payment *Payment, delivery *Delivery, items *[]Item) *Model {
	items_ := make([]ModelItem, len(*items))
	for i, item := range *items {
		items_[i] = ModelItem{
			item.ChrtId,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmId,
			item.Brand,
			item.Status,
		}
	}
	model := Model{
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		ModelDelivery(*delivery),
		ModelPayment(*payment),
		items_,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.Shardkey,
		order.SmId,
		order.DateCreated,
		order.OofShard,
	}
	return &model
}
