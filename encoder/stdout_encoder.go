package encoder

import "github.com/csizsek/prism/entity"

type StdoutEncoder struct {
	input  chan *entity.CommonEntity
	output chan *entity.StdoutEntity
}

func (this *StdoutEncoder) Encode() {
	for {
		commonEntity := <-this.input
		stdoutEntity := entity.NewStdoutEntity()
		stdoutEntity.Msg = commonEntity.StdoutEntity.Msg
		this.output <- stdoutEntity
	}
}

func NewStdoutEncoder(input chan *entity.CommonEntity, output chan *entity.StdoutEntity) *StdoutEncoder {
	encoder := new(StdoutEncoder)
	encoder.input = input
	encoder.output = output
	return encoder
}
