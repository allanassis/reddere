package storages

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
	"time"
)

type Database struct {
	client mongo.Client
	uri    string
}

type Storage interface {
	Connect() error
	Healthcheck() error
}

func (db *Database) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.uri))
	if err != nil {
		return err
	}

	db.client = *client
	return nil
}

func (db *Database) Healthcheck() error {
	err := db.client.Ping(context.Background(), &readpref.ReadPref{})
	if err != nil {
		return err
	}
	return nil
}

func NewDatabase(uri string) Storage {
	db := &Database{
		uri: uri,
	}
	return db
}
