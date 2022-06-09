package services

import "github.com/allanassis/reddere/src/storages"

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
	print("Document inserted with id %s", id)
	return id, nil
}
