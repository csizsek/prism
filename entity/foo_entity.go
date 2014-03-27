package entity

type FooEntity struct {
	Data string
}

func NewFooEntity() *FooEntity {
	entity := new(FooEntity)
	return entity
}
