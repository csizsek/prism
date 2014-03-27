package encoder

import "github.com/csizsek/prism/entity"

type BarEncoder struct {
	input  chan *entity.CommonEntity
	output chan *entity.BarEntity
}

func (this *BarEncoder) Encode() {
	for {
		commonEntity := <-this.input
		this.output <- entity.NewBarEntity(commonEntity.ScribeCategory + " " + commonEntity.ScribeMessage)
	}
}

func NewBarEncoder(input chan *entity.CommonEntity, output chan *entity.BarEntity) *BarEncoder {
	encoder := new(BarEncoder)
	encoder.input = input
	encoder.output = output
	return encoder
}
