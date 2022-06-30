package storages

import (
	"github.com/allanassis/reddere/src/observability"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
	"time"
)

type Database struct {
	client mongo.Client
	uri    string
	logger *observability.Logger
}

type Storage interface {
	Connect() error
	Healthcheck() error
	Bind(result *mongo.SingleResult, instance interface{}) error
	Save(document interface{}, collectionName string) (string, error)
	Get(id string, collectionName string) (*mongo.SingleResult, error)
}

func (db *Database) Save(document interface{}, collectionName string) (string, error) {
	collection := db.client.Database("reddere").Collection(collectionName)

	ctxInsert, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	insertedResult, err := collection.InsertOne(ctxInsert, document)
	if err != nil {
		panic(err)
	}

	stringObjectID := insertedResult.InsertedID.(primitive.ObjectID).Hex()

	return stringObjectID, nil

}

func (db *Database) Get(id string, collectionName string) (*mongo.SingleResult, error) {
	collection := db.client.Database("reddere").Collection(collectionName)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	result := collection.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: objectId}})

	return result, nil
}

func (db *Database) Bind(result *mongo.SingleResult, instance interface{}) error {
	err := result.Decode(instance)
	if err != nil {
		panic(err)
	}
	return nil
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

func NewDatabase(uri string, logger *observability.Logger) Storage {
	db := &Database{
		uri:    uri,
		logger: logger,
	}
	return db
}
