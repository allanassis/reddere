package services

type Template struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}
