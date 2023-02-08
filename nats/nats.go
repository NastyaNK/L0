package nats

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func Publish(b []byte) {
	conn, _ := GetStream()
	conn.Publish(channel, b)
}

var con *nats.Conn //объект подключения к серверу
var channel string

func GetStream() (*nats.Conn, error) {
	var err error
	if con == nil {
		con, err = nats.Connect(nats.DefaultURL) //подключение к серверу
		b, err2 := os.ReadFile("channel")
		if err2 != nil {
			log.Fatal(err)
		}
		channel = string(b)
	}
	return con, err
}
