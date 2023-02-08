package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

type NATS struct {
	con     *nats.Conn
	channel string
}
type getter func([]byte)

func (n *NATS) Publish(b []byte) {
	n.con.Publish(n.channel, b)
}
func (n *NATS) Connect(callback getter) {
	nc, err := n.GetStream()
	if err != nil {
		log.Fatal(err)
	}
	nc.Subscribe(n.channel, func(msg *nats.Msg) {
		fmt.Println("Получен файл")
		callback(msg.Data)
	})
}

func (n *NATS) GetStream() (*nats.Conn, error) {
	var err error
	if n.con == nil {
		n.con, err = nats.Connect(nats.DefaultURL) //подключение к серверу
		b, err2 := os.ReadFile("channel")
		if err2 != nil {
			log.Fatal(err)
		}
		n.channel = string(b)
	}
	return n.con, err
}
