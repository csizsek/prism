package entity

type FooEntity struct {
	data string
}

func NewFooEntity(data string) *FooEntity {
	fmt.Println("NewFooEntity")
	e := new(FooEntity)
	e.data = data
	return e
}
