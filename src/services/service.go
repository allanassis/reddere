package services

import (
	"fmt"

	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/storages"
)

type Service interface {
	Save(entity interface{}, entityName string) (string, error)
	Build(entity interface{}, entityID string, entityName string) error
	Delete(entityID string) error
}

type BaseService struct {
	storage storages.Storage
	logger  *logging.Logger
}

func (service *BaseService) Save(entity interface{}, entityName string) (string, error) {
	id, err := service.storage.Save(entity, entityName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with id %s\n", id)
	return id, nil
}

func (service *BaseService) Build(entity interface{}, entityID string, entityName string) error {
	result, err := service.storage.Get(entityID, entityName)
	if err != nil {
		panic(err)
	}
	service.storage.Bind(result, entity)
	fmt.Printf("Document retrieved %+v\n", entity)
	return nil
}

func (service *BaseService) Delete(entityID string) error {
	_, err := service.storage.Delete(entityID, "template")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document Deleted %+v\n", entityID)
	return nil
}

// type ServiceResult interface {

// }

func NewService(storage storages.Storage, logger *logging.Logger) Service {
	return &BaseService{
		storage: storage,
		logger:  logger,
	}
}
