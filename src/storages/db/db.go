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
		logError("Creating mongo client error", err, logger, db)
		panic(err)
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

	db.logger.Debug("Retrive collection from database", logging.String("collection", collectionName))

	insertedResult, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		logError("when inserting document into database", err, logger, db)
		return "", err
	}

	stringObjectID := insertedResult.InsertedID.(primitive.ObjectID).Hex()

	db.logger.Debug("Succefully inserted document into database", logging.String("id", stringObjectID), logging.String("collection", collectionName))
	return stringObjectID, nil
}

func (db *Database) Get(id string, collectionName string) (*mongo.SingleResult, error) {
	collection := db.instance.Collection(collectionName)
	logger := db.logger.With(logging.String("collection", collectionName))
	logger.Debug("Retrive collection from database")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logError("invalid id when get item from database", err, logger, db)
		return nil, err
	}
	logger.Debug("Succefully parsed id into object id")

	result := collection.FindOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: objectId}})
	if result.Err() == mongo.ErrNoDocuments {
		logger.Warn("No document found in database", logging.String("id", id))
		return nil, nil
	}

	if result.Err() != nil {
		logError("when try to find document in database", err, logger, db)
		return nil, err
	}

	logger.Debug("Succefully found item in database", logging.String("id", id))
	return result, nil
}

func (db *Database) Delete(id string, collectionName string) (*mongo.DeleteResult, error) {
	collection := db.instance.Collection(collectionName)
	logger := db.logger.With(logging.String("collection", collectionName))

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logError("invalid id when deleting item from database", err, logger, db)
		return nil, err
	}

	result, err := collection.DeleteOne(context.Background(), bson.D{primitive.E{Key: "_id", Value: objectId}})
	if err != nil {
		logError("when deleting item in database", err, logger, db)
		return nil, err
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
		logError("when binding result to instance", err, db.logger, db)
		panic(err)
	}

	db.logger.Debug("Succefully bind item", logging.Any("item", instance))
	return nil
}

func (db *Database) Healthcheck() error {
	err := db.client.Ping(context.Background(), &readpref.ReadPref{})
	if err != nil {
		logError("when pinging to database", err, db.logger, db)
		return err
	}

	db.logger.Info("Succefully ping to database")
	return nil
}
