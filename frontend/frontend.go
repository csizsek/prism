package frontend

import (
	"github.com/csizsek/prism/config"
	"github.com/csizsek/prism/decoder"
	"github.com/csizsek/prism/entity"
	"github.com/csizsek/prism/receiver"
)

type Frontend interface {
	Start()
}

func NewFrontend(config config.FrontendConfig, dec2disp chan *entity.CommonEntity) *Frontend {
	var ret Frontend
	switch config.Protocol {
	case "http":
		frontend := *NewHttpFrontend()
		frontend.rec2dec = make(chan *entity.HttpEntity)
		frontend.dec2disp = dec2disp
		frontend.receiver = receiver.NewHttpReceiver(frontend.rec2dec)
		frontend.receiver.Configure(config.HttpReceiverConfig)
		frontend.decoder = decoder.NewHttpDecoder(frontend.rec2dec, frontend.dec2disp)
		ret = frontend
	}
	return &ret
}
