package entities

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Template struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name"`
	Url  string `json:"url,omitempty" bson:"url"`
}

func (template *Template) EntityName() string {
	return "template"
}

func (template *Template) IsValid() (bool, error) {
	err := validation.ValidateStruct(
		template,
		validation.Field(&template.Name, validation.Required),
		validation.Field(&template.Name, is.Alpha),
		validation.Field(&template.Name, validation.Length(2, 50)),
		// Validate if this URL already exists in database
		validation.Field(&template.Url, validation.Required),
		validation.Field(&template.Url, is.URL),
		validation.Field(&template.Url, validation.Length(5, 500)),
	)

	if err != nil {
		return false, err
	}

	return true, nil
}
