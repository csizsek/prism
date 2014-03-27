package entity

type CommonEntity struct {
	FooData        string
	BarMsg         string
	ScribeCategory string
	ScribeMessage  string
}

func NewCommonEntity() *CommonEntity {
	entity := new(CommonEntity)
	return entity
}
