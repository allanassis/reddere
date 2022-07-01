package storages

import "go.mongodb.org/mongo-driver/mongo"

type Storage interface {
	Healthcheck() error
	Bind(result *mongo.SingleResult, instance interface{}) error
	Save(document interface{}, collectionName string) (string, error)
	Get(id string, collectionName string) (*mongo.SingleResult, error)
	Delete(id string, collectionName string) (*mongo.DeleteResult, error)
}
