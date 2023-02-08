package nats_sender

import (
	"L0/internal/entity"
	"L0/internal/nats"
	"L0/internal/rand"
	"encoding/json"
	"log"
	"os"
	"time"
)

func Send(n *nats.NATS, count int, interval time.Duration) {
	mod := GetModel("model.json")
	for i := 0; i < count; i++ {
		time.Sleep(interval * time.Second)
		mod.OrderUid = rand.GenerateString(rand.LetterBytes+rand.NumberBytes, len(mod.OrderUid))
		mod.Delivery.Phone = "+7" + rand.GenerateString(rand.NumberBytes, 10)
		mod.Payment.Transaction = rand.GenerateString(rand.LetterBytes+rand.NumberBytes, len(mod.Payment.Transaction))
		for i := 0; i < len(mod.Items); i++ {
			mod.Items[i].Rid = rand.GenerateString(rand.LetterBytes+rand.NumberBytes, len(mod.Items[i].Rid))
			mod.Items[i].ChrtId = rand.GenerateNumber(7)
			mod.Items[i].NmId = rand.GenerateNumber(7)
		}
		n.Publish(ToBytes(mod))
	}
}

func GetModel(str string) *entity.Model {
	data, err := os.ReadFile(str)
	if err != nil {
		log.Fatal(err)
	}
	var variable entity.Model
	err = json.Unmarshal(data, &variable)
	if err != nil {
		log.Fatal(err)
	}
	return &variable
}
func ToBytes(model *entity.Model) []byte {
	b, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
