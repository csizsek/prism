package decoder

import "github.com/csizsek/prism/entity"

type ScribeDecoder struct {
	input  chan *entity.ScribeEntity
	output chan *entity.CommonEntity
}

func (this *ScribeDecoder) Decode() {
	for {
		scribeEntity := <-this.input
		commonEntity := entity.NewCommonEntity()
		commonEntity.ScribeEntity.Category = scribeEntity.Category
		commonEntity.ScribeEntity.Message = scribeEntity.Message
		this.output <- commonEntity
	}
}

func NewScribeDecoder(input chan *entity.ScribeEntity, output chan *entity.CommonEntity) *ScribeDecoder {
	decoder := new(ScribeDecoder)
	decoder.input = input
	decoder.output = output
	return decoder
}
