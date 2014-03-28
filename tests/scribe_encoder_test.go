package main

import "testing"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/encoder"

func TestScribeEncoderEncode(t *testing.T) {
	input := make(chan *entity.CommonEntity)
	output := make(chan *entity.ScribeEntity)

	scribeEncoder := encoder.NewScribeEncoder(input, output)
	go scribeEncoder.Encode()

	commonEntity := entity.NewCommonEntity()
	commonEntity.ScribeEntity.Category = "test category"
	commonEntity.ScribeEntity.Message = "test message"
	input <- commonEntity
	scribeEntity := <-output

	if scribeEntity.Category != "test category" {
		t.Fail()
	}
	if scribeEntity.Message != "test message" {
		t.Fail()
	}
}
