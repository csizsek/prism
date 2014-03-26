package sender

import "fmt"
import "github.com/csizsek/prism/entity"

type BarSender struct {
	input chan entity.BarEntity
}

func (this *BarSender) Send() {
	for {
		barEntity := <-this.input
		fmt.Println(barEntity.Msg)
	}
}

func NewBarSender(input chan entity.BarEntity) *BarSender {
	sender := new(BarSender)
	sender.input = input
	return sender
}
