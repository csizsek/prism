package backend

import (
	"github.com/csizsek/prism/encoder"
	"github.com/csizsek/prism/entity"
	"github.com/csizsek/prism/mapper"
	"github.com/csizsek/prism/sender"
)

type StdoutBackend struct {
	disp2map chan *entity.CommonEntity
	map2enc  chan *entity.CommonEntity
	enc2send chan *entity.StdoutEntity
	mapper   *mapper.Mapper
	encoder  *encoder.StdoutEncoder
	sender   *sender.StdoutSender
}

func (this StdoutBackend) Start() {
	go this.sender.Send()
	go this.encoder.Encode()
	go this.mapper.Map()
}

func NewStdoutBackend() *StdoutBackend {
	return new(StdoutBackend)
}
