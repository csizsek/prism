package main

import "testing"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/dispatcher"

func TestDispatch(t *testing.T) {
	input := make(chan *entity.CommonEntity)
	outputs := []chan *entity.CommonEntity{make(chan *entity.CommonEntity), make(chan *entity.CommonEntity)}

	dispatcher := dispatcher.NewDispatcher(input, outputs)
	go dispatcher.Dispatch()

	commonEntity1 := entity.NewCommonEntity()
	commonEntity1.BarEntity.Msg = "test msg"
	input <- commonEntity1

	commonEntity2 := <-outputs[0]
	if commonEntity2.BarEntity.Msg != "test msg" {
		t.Fail()
	}

	commonEntity3 := <-outputs[1]
	if commonEntity3.BarEntity.Msg != "test msg" {
		t.Fail()
	}
}
