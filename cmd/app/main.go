package main

import (
	"L0/cmd/nats_sender"
	"L0/internal/app"
	"L0/internal/entity"
	"L0/internal/web"
	"os"
)

func main() {
	app_ := app.NewAPP(os.Args[1], os.Args[2])
	web := web.Web{
		Files: "./front",
		GetModel: func(id string) *entity.Model {
			return app_.Read(id)
		},
	}
	go nats_sender.Send(&app_.NATS, 100, 1)
	web.Run(os.Args[3])
}
