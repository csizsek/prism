package entity

type BarEntity struct {
	Msg string
}

func NewBarEntity(msg string) *BarEntity {
	entity := new(BarEntity)
	entity.Msg = msg
	return entity
}
