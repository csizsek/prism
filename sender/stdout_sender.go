package sender

import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/config"

type StdoutSender struct {
	input  chan *entity.StdoutEntity
	prefix string
}

func (this *StdoutSender) Send() {
	for {
		stdoutEntity := <-this.input
		println(this.prefix, stdoutEntity.Msg)
	}
}

func (this *StdoutSender) Configure(config config.StdoutSenderConfig) {
	this.prefix = config.Prefix
}

func NewStdoutSender(input chan *entity.StdoutEntity) *StdoutSender {
	sender := new(StdoutSender)
	sender.input = input
	return sender
}
