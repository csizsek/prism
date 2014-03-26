package entity

type BarEntity struct {
	msg string
}

func NewBarEntity(msg string) *BarEntity {
	fmt.Println("NewBarEntity")
	e := new(BarEntity)
	e.msg = msg
	return e
}