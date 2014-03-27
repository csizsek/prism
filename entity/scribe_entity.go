package entity

type ScribeEntity struct {
	Category string
	Message  string
}

func NewScribeEntity(category, message string) *ScribeEntity {
	entity := new(ScribeEntity)
	entity.Category = category
	entity.Message = message
	return entity
}
