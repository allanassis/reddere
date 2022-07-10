package entities

type Template struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name"`
	Url  string `json:"url,omitempty" bson:"url"`
}

func (template *Template) EntityName() string {
	return "template"
}
