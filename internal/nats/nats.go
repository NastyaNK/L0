package nats

import (
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
func (n *NATS) Connect(file string, callback getter) {
	nc, err := n.GetStream(file)
	if err != nil {
		log.Fatal(err)
	}
	nc.Subscribe(n.channel, func(msg *nats.Msg) {
		callback(msg.Data)
	})
}

func (n *NATS) GetStream(file string) (*nats.Conn, error) {
	var err error
	if n.con == nil {
		n.con, err = nats.Connect(nats.DefaultURL) //подключение к серверу
		b, err2 := os.ReadFile(file)
		if err2 != nil {
			log.Fatal(err)
		}
		n.channel = string(b)
	}
	return n.con, err
}
