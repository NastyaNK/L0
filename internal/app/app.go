//основная логика сервиса

package app

import (
	"L0/internal/cache"
	"L0/internal/entity"
	"L0/internal/nats"
	psql "L0/internal/repository"
	"encoding/json"
	"fmt"
	"log"
)

type App struct {
	PDB   *psql.Postgres
	NATS  nats.NATS
	cache *cache.Cache
}

func NewAPP(config, channel string) *App {
	var app = App{
		PDB:   psql.NewDB(psql.GetDBConfig(config)),
		NATS:  nats.NATS{},
		cache: cache.NewCache(),
	}
	models := app.PDB.GetAll()
	fmt.Println("Загружено в кэш:", len(models))
	app.cache.SetAll(models)
	app.NATS.Connect(channel, func(b []byte) {
		var model = entity.Model{}
		err := json.Unmarshal(b, &model)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Получен новый заказ:", model.OrderUid)
		app.cache.Set(&model)
		app.PDB.Set(&model)
	})
	return &app
}

func (a *App) Read(id string) *entity.Model {
	return a.cache.Get(id)
}
