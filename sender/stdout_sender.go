package sender

import "fmt"
import "github.com/csizsek/prism/entity"

type StdoutSender struct {
	input chan *entity.StdoutEntity
}

func (this *StdoutSender) Send() {
	for {
		stdoutEntity := <-this.input
		fmt.Println(stdoutEntity.Msg)
	}
}

func NewStdoutSender(input chan *entity.StdoutEntity) *StdoutSender {
	sender := new(StdoutSender)
	sender.input = input
	return sender
}
