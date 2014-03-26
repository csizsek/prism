package entity

type CommonEntity struct {
	foo_data string
	bar_msg string
}

func NewCommonEntity(foo_data string, bar_msg string) *CommonEntity {
	fmt.Println("NewCommonEntity")
	e := new(CommonEntity)
	e.foo_data = foo_data
	e.bar_msg = bar_msg
	return e
}
