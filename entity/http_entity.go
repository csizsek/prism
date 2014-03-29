package entity

type HttpEntity struct {
	Data string
}

func NewHttpEntity() *HttpEntity {
	return new(HttpEntity)
}
