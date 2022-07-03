package services

import (
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services/entities"
	"github.com/allanassis/reddere/src/storages"
)

type Service interface {
	Save(entity entities.Entity) (string, error)
	Build(entity entities.Entity, entityID string) error
	Delete(entityID string, entityName string) error
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
	return id, nil
}

func (service *BaseService) Build(entity entities.Entity, entityID string) error {
	result, err := service.storage.Get(entityID, entity.EntityName())
	if err != nil {
		panic(err)
	}
	service.storage.Bind(result, entity)
	return nil
}

func (service *BaseService) Delete(entityID string, entityName string) error {
	_, err := service.storage.Delete(entityID, entityName)
	if err != nil {
		panic(err)
	}
	return nil
}

// type ServiceResult interface {

// }
