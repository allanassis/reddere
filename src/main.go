package main

import (
	"context"
	"log"
	"time"

	"github.com/allanassis/reddere/src/api"
	"github.com/allanassis/reddere/src/config"
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services"
	"github.com/allanassis/reddere/src/storages/db"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			services.NewService,
			config.NewConfig,
			logging.NewLogger,
			db.NewDatabase,
		),
		fx.Invoke(api.NewServer),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
}
