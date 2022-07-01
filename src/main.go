package main

import (
	"context"
	"log"
	"time"

	"github.com/allanassis/reddere/src/api"
	"github.com/allanassis/reddere/src/config"
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services"
	"github.com/allanassis/reddere/src/storages"
	"go.uber.org/fx"
)

func Register(service services.Service, db storages.Storage, logger *logging.Logger) {
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	api.InitServer(service, db, logger)
}

func main() {
	app := fx.New(
		fx.Provide(
			services.NewService,
			config.NewConfig,
			logging.NewLogger,
			storages.NewDatabase,
		),
		fx.Invoke(Register),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
}
