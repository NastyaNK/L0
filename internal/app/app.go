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

func NewAPP(config, channel string) (*App, error) {
	dbConfig, err := psql.GetDBConfig(config)
	if err != nil {
		return nil, err
	}
	db, err := psql.NewDB(dbConfig)
	if err != nil {
		return nil, err
	}
	var app = App{
		PDB:   db,
		NATS:  nats.NATS{},
		cache: cache.NewCache(),
	}
	models, err := app.PDB.GetAll()
	if err == nil {
		app.cache.SetAll(models)
	}

	err = app.NATS.Connect(channel, func(b []byte) {
		var model = entity.Model{}
		err := json.Unmarshal(b, &model)
		if err != nil {
			log.Fatal(err)
		}
		app.cache.Set(&model)
		err = app.PDB.Set(&model)
		if err != nil {
			log.Fatal(err)
		}
	})
	return &app, err
}

func (a *App) Read(id string) *entity.Model {
	return a.cache.Get(id)
}
