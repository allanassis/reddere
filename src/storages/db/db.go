package db

import (
	"time"

	"github.com/allanassis/reddere/src/config"
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/storages"
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
	timeout  time.Duration
}

func NewDatabase(logger *logging.Logger, config *config.Config) storages.Storage {
	uri := config.GetString("database.uri")
	name := config.GetString("database.name")
	timeout := config.GetDuration("database.timeout")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	instance := client.Database(name)

	db := &Database{
		client:   *client,
		instance: instance,
		logger:   logger,
		timeout:  timeout,
	}
	if err != nil {
		panic(newDbError(DB_CREATE_CLIENT_ERROR, err))
	}

	err = db.Healthcheck()
	if err != nil {
		panic(err)
	}

	return db
}

func (db *Database) Save(document interface{}, collectionName string) (string, error) {
	collection := db.instance.Collection(collectionName)
	logger := db.logger.With(logging.String("collection", collectionName))

	logger.Debug("Retrive collection from database", logging.String("collection", collectionName))

	insertedResult, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return "", newDbError(DB_INSER_ONE_ERROR, err)
	}

	stringObjectID := insertedResult.InsertedID.(primitive.ObjectID).Hex()

	logger.Debug("Succefully inserted document into database", logging.String("id", stringObjectID), logging.String("collection", collectionName))
	return stringObjectID, nil
}

func (db *Database) Get(id string, collectionName string) (*mongo.SingleResult, error) {
	collection := db.instance.Collection(collectionName)
	logger := db.logger.With(logging.String("collection", collectionName))
	logger.Debug("Retrive collection from database")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, newDbError(DB_INVALID_ID_ERROR, err)
	}
	logger.Debug("Succefully parsed id into object id")

	result := collection.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: objectId}})
	if result.Err() == mongo.ErrNoDocuments {
		logger.Warn("No document found in database", logging.String("id", id))
		return nil, nil
	}

	if result.Err() != nil {
		return nil, newDbError(DB_FIND_ONE_ERROR, err)
	}

	logger.Debug("Succefully found item in database", logging.String("id", id))
	return result, nil
}

func (db *Database) Delete(id string, collectionName string) (*mongo.DeleteResult, error) {
	collection := db.instance.Collection(collectionName)
	logger := db.logger.With(logging.String("collection", collectionName))

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, newDbError(DB_INVALID_ID_ERROR, err)
	}

	result, err := collection.DeleteOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: objectId}})
	if err != nil {
		return nil, newDbError(DB_DELETE_ONE_ERROR, err)
	}

	logger.Debug("Succefully deleted item in database", logging.String("id", id))
	return result, nil
}

func (db *Database) Bind(result *mongo.SingleResult, instance interface{}) error {
	err := result.Decode(instance)
	if err == mongo.ErrNoDocuments {
		db.logger.Warn("No document to bind")
		return nil
	}

	if err != nil {
		return newDbError(DB_BIND_ERROR, err)
	}

	db.logger.Debug("Succefully bind item", logging.Any("item", instance))
	return nil
}

func (db *Database) Healthcheck() error {
	err := db.client.Ping(context.Background(), &readpref.ReadPref{})
	if err != nil {
		return newDbError(DB_HEALTHCHECK_ERROR, err)
	}

	db.logger.Info("Succefully ping to database")
	return nil
}
