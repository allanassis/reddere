package services

import (
	"fmt"

	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services/entities"
	"github.com/allanassis/reddere/src/storages"
)

type Service interface {
	Save(entity entities.Entity) (string, error)
	Build(entity entities.Entity, entityID string) error
	Delete(entityID string) error
}

type BaseService struct {
	storage storages.Storage
	logger  *logging.Logger
}

func NewService(storage storages.Storage, logger *logging.Logger) Service {
	return &BaseService{
		storage: storage,
		logger:  logger,
	}
}

func (service *BaseService) Save(entity entities.Entity) (string, error) {
	id, err := service.storage.Save(entity, entity.EntityName())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with id %s\n", id)
	return id, nil
}

func (service *BaseService) Build(entity entities.Entity, entityID string) error {
	result, err := service.storage.Get(entityID, entity.EntityName())
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
