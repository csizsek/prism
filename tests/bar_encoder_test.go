package main

import "testing"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/encoder"

func TestBarEncoderEncode(t *testing.T) {
	input := make(chan *entity.CommonEntity)
	output := make(chan *entity.BarEntity)

	barEncoder := encoder.NewBarEncoder(input, output)
	go barEncoder.Encode()
	
	commonEntity := entity.NewCommonEntity()
	commonEntity.BarEntity.Msg = "test data"
	input <-commonEntity 
	barEntity := <-output
	
	if barEntity.Msg != "test data" {
		t.Fail()
	}
}
