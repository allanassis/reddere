package main

import (
	"github.com/allanassis/reddere/src/api"
	"github.com/allanassis/reddere/src/observability"
	"github.com/allanassis/reddere/src/storages"
)

func main() {
	logger := observability.Default()

	db := storages.NewDatabase("mongodb://localhost:27017", logger)
	db.Connect()
	api.InitServer(db, logger)

}
