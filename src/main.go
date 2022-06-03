package main

import (
	"github.com/allanassis/reddere/src/storages"
)

func main() {
	db := storages.NewDatabase("mongodb://localhost:27017")
	db.Connect()
	err := db.Healthcheck()
	if err != nil {
		panic(err)
	}
	print("Connected do db :D")
}
