package main

import (
	"github.com/allanassis/reddere/src/api"
	"github.com/allanassis/reddere/src/storages"
)

func main() {
	db := storages.NewDatabase("mongodb://localhost:27017")
	db.Connect()
	api.InitServer(db)
}
