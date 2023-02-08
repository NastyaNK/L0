//основная логика сервиса

package app

import (
	"L0/internal/cache"
	"L0/internal/entity"
	"L0/internal/nats"
	psql "L0/internal/repository"
	"encoding/json"
	"log"
)

type App struct {
	PDB   *psql.Postgres
	NATS  nats.NATS
	cache *cache.Cache
}

func NewAPP() *App {
	var app = App{
		PDB:   psql.NewDB("anastasia", "2553", "db", "localhost"),
		NATS:  nats.NATS{},
		cache: cache.NewCache(),
	}
	models := app.PDB.GetAll()
	app.cache.SetAll(models)
	app.NATS.Connect(func(b []byte) {
		var model = entity.Model{}
		err := json.Unmarshal(b, &model)
		if err != nil {
			log.Fatal(err)
		}
		app.cache.Set(&model)
		app.PDB.Set(&model)
	})
	return &app
}

func (a *App) Read(id string) *entity.Model {
	return a.cache.Get(id)
}
