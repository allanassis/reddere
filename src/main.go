package main

import (
	"context"
	"log"
	"time"

	"github.com/allanassis/reddere/src/api"
	"github.com/allanassis/reddere/src/observability"
	"github.com/allanassis/reddere/src/storages"
	"go.uber.org/fx"
)

func Register(db storages.Storage, logger *observability.Logger) {
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	api.InitServer(db, logger)
}

func main() {
	app := fx.New(
		fx.Provide(
			observability.Default,
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
