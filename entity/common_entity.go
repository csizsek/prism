package entity

type CommonEntity struct {
	HttpEntity
	StdoutEntity
	ScribeEntity
}

func NewCommonEntity() *CommonEntity {
	return new(CommonEntity)
}
