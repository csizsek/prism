package entity

type ScribeEntity struct {
	Category string
	Message  string
}

func NewScribeEntity() *ScribeEntity {
	return new(ScribeEntity)
}
