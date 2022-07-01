package entities

type Template struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

func (template *Template) EntityName() string {
	return "template"
}
