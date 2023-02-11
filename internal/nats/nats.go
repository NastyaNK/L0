package nats

import (
	"github.com/nats-io/nats.go"
	"os"
)

type NATS struct {
	con     *nats.Conn
	channel string
}
type getter func([]byte)

func (n *NATS) Publish(b []byte) error {
	err := n.con.Publish(n.channel, b)
	return err
}
func (n *NATS) Connect(file string, callback getter) error {
	nc, err := n.GetStream(file)
	if err != nil {
		return err
	}
	_, err = nc.Subscribe(n.channel, func(msg *nats.Msg) {
		callback(msg.Data)
	})
	return err
}

func (n *NATS) GetStream(file string) (*nats.Conn, error) {
	var err error
	if n.con == nil {
		n.con, err = nats.Connect(nats.DefaultURL)
		b, err2 := os.ReadFile(file)
		if err2 != nil {
			return nil, err2
		}
		n.channel = string(b)
	}
	return n.con, err
}
