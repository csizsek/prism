package dispatcher

import "github.com/csizsek/prism/entity"

type Dispatcher struct {
	input  chan *entity.CommonEntity
	output []chan *entity.CommonEntity
}

func (this *Dispatcher) Dispatch() {
	for {
		entity := <-this.input
		for _, channel := range this.output {
			channel <- entity
		}
	}
}

func NewDispatcher(input chan *entity.CommonEntity, output []chan *entity.CommonEntity) *Dispatcher {
	dispatcher := new(Dispatcher)
	dispatcher.input = input
	dispatcher.output = output
	return dispatcher
}
