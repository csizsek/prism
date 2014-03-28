package encoder

import "github.com/csizsek/prism/entity"

type ScribeEncoder struct {
	input  chan *entity.CommonEntity
	output chan *entity.ScribeEntity
}

func (this *ScribeEncoder) Encode() {
	for {
		commonEntity := <-this.input
		scribeEntity := entity.NewScribeEntity()
		scribeEntity.Category = commonEntity.ScribeEntity.Category
		scribeEntity.Message = commonEntity.ScribeEntity.Message
		this.output <- scribeEntity
	}
}

func NewScribeEncoder(input chan *entity.CommonEntity, output chan *entity.ScribeEntity) *ScribeEncoder {
	encoder := new(ScribeEncoder)
	encoder.input = input
	encoder.output = output
	return encoder
}
