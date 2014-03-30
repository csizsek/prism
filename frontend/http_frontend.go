package frontend

import (
	"github.com/csizsek/prism/decoder"
	"github.com/csizsek/prism/entity"
	"github.com/csizsek/prism/receiver"
)

type HttpFrontend struct {
	rec2dec  chan *entity.HttpEntity
	dec2disp chan *entity.CommonEntity
	receiver *receiver.HttpReceiver
	decoder  *decoder.HttpDecoder
}

func (this HttpFrontend) Start() {
	go this.decoder.Decode()
	go this.receiver.Receive()
}

func NewHttpFrontend() *HttpFrontend {
	return new(HttpFrontend)
}
