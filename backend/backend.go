package backend

import (
	"github.com/csizsek/prism/config"
	"github.com/csizsek/prism/encoder"
	"github.com/csizsek/prism/entity"
	"github.com/csizsek/prism/mapper"
	"github.com/csizsek/prism/sender"
)

type Backend interface {
	Start()
}

func NewBackend(config config.BackendConfig, disp2map chan *entity.CommonEntity) *Backend {
	var ret Backend
	switch config.Protocol {
	case "stdout":
		backend := *NewStdoutBackend()
		backend.disp2map = disp2map
		backend.map2enc = make(chan *entity.CommonEntity)
		backend.enc2send = make(chan *entity.StdoutEntity)
		backend.mapper = mapper.NewMapper(backend.disp2map, backend.map2enc)
		backend.encoder = encoder.NewStdoutEncoder(backend.map2enc, backend.enc2send)
		backend.sender = sender.NewStdoutSender(backend.enc2send)
		backend.sender.Configure(config.StdoutSenderConfig)
		ret = backend
	}
	return &ret
}
