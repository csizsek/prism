package entity

type ScribeEntity struct {
	Category string
	Message  string
}

func NewScribeEntity() *ScribeEntity {
	entity := new(ScribeEntity)
	return entity
}
