package decoder

import "github.com/csizsek/prism/entity"

type FooDecoder struct {
	input chan entity.FooEntity
	output chan entity.CommonEntity
}

func (this *FooDecoder) Decode() {
	for {
		fooEntity := <- this.input
		this.output <- *entity.NewCommonEntity(fooEntity.Data, "")
	}
}

func NewFooDecoder(input chan entity.FooEntity, output chan entity.CommonEntity) *FooDecoder {
	decoder := new(FooDecoder)
	decoder.input = input
	decoder.output = output
	return decoder
}
