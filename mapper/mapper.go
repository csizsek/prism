package mapper

import "github.com/csizsek/prism/entity"

type Mapper struct {
	input  chan entity.CommonEntity
	output chan entity.CommonEntity
}

func (this *Mapper) Map() {
	for {
		entity := <-this.input
		this.output <-entity
	}
}

func NewMapper(input chan entity.CommonEntity, output chan entity.CommonEntity) *Mapper {
	mapper := new(Mapper)
	mapper.input = input
	mapper.output = output
	return mapper
}

