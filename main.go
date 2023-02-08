package main

//отправщик в nats-sremeng
import (
	"L0/entity"
	"L0/nats"
	"L0/rand"
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {
	mod := GetModel("model.json")
	for {
		time.Sleep(time.Second)
		mod.OrderUid = rand.GenerateString(rand.LetterBytes+rand.NumberBytes, len(mod.OrderUid))
		mod.Delivery.Phone = "+7" + rand.GenerateString(rand.NumberBytes, 10)
		mod.Payment.Transaction = rand.GenerateString(rand.LetterBytes+rand.NumberBytes, len(mod.Payment.Transaction))
		for i := 0; i < len(mod.Items); i++ {
			mod.Items[i].Rid = rand.GenerateString(rand.LetterBytes+rand.NumberBytes, len(mod.Items[i].Rid))
			mod.Items[i].ChrtId = rand.GenerateNumber(7)
			mod.Items[i].NmId = rand.GenerateNumber(7)
		}
		nats.Publish(ToBytes(mod))
	}
}

func GetModel(str string) *entity.Model {
	data, err := os.ReadFile(str)
	if err != nil {
		log.Fatal(err)
	}
	var variable entity.Model
	err = json.Unmarshal(data, &variable) //из байтов в структурy передача
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
