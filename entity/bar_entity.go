package entity

type BarEntity struct {
	Msg string
}

func NewBarEntity() *BarEntity {
	entity := new(BarEntity)
	return entity
}
