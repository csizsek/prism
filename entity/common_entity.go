package entity

type CommonEntity struct {
	FooEntity
	BarEntity
	ScribeEntity
}

func NewCommonEntity() *CommonEntity {
	entity := new(CommonEntity)
	return entity
}
