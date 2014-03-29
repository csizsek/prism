package decoder

import "github.com/csizsek/prism/entity"

type HttpDecoder struct {
	input  chan *entity.HttpEntity
	output chan *entity.CommonEntity
}

func (this *HttpDecoder) Decode() {
	for {
		httpEntity := <-this.input
		commonEntity := entity.NewCommonEntity()
		commonEntity.HttpEntity.Data = httpEntity.Data
		this.output <- commonEntity
	}
}

func NewHttpDecoder(input chan *entity.HttpEntity, output chan *entity.CommonEntity) *HttpDecoder {
	decoder := new(HttpDecoder)
	decoder.input = input
	decoder.output = output
	return decoder
}
