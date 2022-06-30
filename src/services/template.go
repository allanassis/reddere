package services

import (
	"fmt"

	"github.com/allanassis/reddere/src/storages"
)

type Template struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

func (template *Template) Save(storage storages.Storage) (string, error) {
	id, err := storage.Save(template, "template")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with id %s\n", id)
	return id, nil
}

func (template *Template) Build(templateId string, storage storages.Storage) error {
	result, err := storage.Get(templateId, "template")
	if err != nil {
		panic(err)
	}
	storage.Bind(result, template)
	fmt.Printf("Document retrieved %+v\n", template)
	return nil
}

func (template *Template) Delete(storage storages.Storage) error {
	_, err := storage.Delete(template.ID, "template")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document Deleted %+v\n", template.ID)
	return nil
}
