package entity

type CommonEntity struct {
	FooData string
	BarMsg string
}

func NewCommonEntity(fooData string, barMsg string) *CommonEntity {
	entity := new(CommonEntity)
	entity.FooData = fooData
	entity.BarMsg = barMsg
	return entity
}
