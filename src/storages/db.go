package storages

import (
	"github.com/allanassis/reddere/src/config"
	"github.com/allanassis/reddere/src/observability/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"context"
)

type Database struct {
	client   mongo.Client
	instance *mongo.Database
	logger   *logging.Logger
}

func NewDatabase(logger *logging.Logger, config *config.Config) Storage {
	uri := config.GetString("database.uri")
	name := config.GetString("database.name")
	timeout := config.GetDuration("database.timeout")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal("Panic when connecting to database",
			logging.String("name", name),
			logging.String("timeout", timeout.String()),
			logging.String("uri", uri))
		panic(err)
	}

	instance := client.Database(name)

	db := &Database{
		client:   *client,
		instance: instance,
		logger:   logger,
	}

	return db
}

func (db *Database) Save(document interface{}, collectionName string) (string, error) {
	collection := db.instance.Collection(collectionName)

	insertedResult, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return "", err
	}

	stringObjectID := insertedResult.InsertedID.(primitive.ObjectID).Hex()

	return stringObjectID, nil

}

func (db *Database) Get(id string, collectionName string) (*mongo.SingleResult, error) {
	collection := db.instance.Collection(collectionName)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := collection.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: objectId}})

	return result, nil
}

func (db *Database) Delete(id string, collectionName string) (*mongo.DeleteResult, error) {
	collection := db.instance.Collection(collectionName)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	result, err := collection.DeleteOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: objectId}})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *Database) Bind(result *mongo.SingleResult, instance interface{}) error {
	err := result.Decode(instance)
	if err != nil {
		panic(err)
	}
	return nil
}

func (db *Database) Healthcheck() error {
	err := db.client.Ping(context.Background(), &readpref.ReadPref{})
	if err != nil {
		return err
	}
	return nil
}
