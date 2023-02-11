package main

import (
	"L0/cmd/nats_sender"
	"L0/internal/app"
	"L0/internal/entity"
	"L0/internal/web"
	"log"
	"os"
)

func main() {
	app_, err := app.NewAPP(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	web := web.Web{
		Files: "./front",
		GetModel: func(id string) *entity.Model {
			return app_.Read(id)
		},
	}
	go nats_sender.Send(&app_.NATS, 100, 1)
	err = web.Run(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
}
