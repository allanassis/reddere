package entities

type Entity interface {
	EntityName() string
	IsValid() (bool, error)
}
