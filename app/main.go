package main

//отправщик в nats-sremeng
import (
	"NATS/app/entity"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=anastasia password=2553 dbname=db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	item := entity.Item{}
	err = db.Get(&item, "SELECT * FROM items")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(item)
	nk()
}

func nk() {
	data, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatal(err)
	}
	var variable entity.Model
	err = json.Unmarshal(data, &variable) //из байтов в структурy передача
	order, payment, del, it := variable.ToDB()
	fmt.Println(order)
	fmt.Println(payment)
	fmt.Println(del)
	fmt.Println(it)
}
