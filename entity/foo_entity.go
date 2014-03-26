package entity

type FooEntity struct {
	Data string
}

func NewFooEntity(data string) *FooEntity {
	entity := new(FooEntity)
	entity.Data = data
	return entity
}
