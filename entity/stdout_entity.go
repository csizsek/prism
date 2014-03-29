package entity

type StdoutEntity struct {
	Msg string
}

func NewStdoutEntity() *StdoutEntity {
	return new(StdoutEntity)
}
