package main

import "testing"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/decoder"

func TestScribeDecoderDecode(t *testing.T) {
	input := make(chan *entity.ScribeEntity)
	output := make(chan *entity.CommonEntity)

	scribeDecoder := decoder.NewScribeDecoder(input, output)
	go scribeDecoder.Decode()

	scribeEntity := entity.NewScribeEntity()
	scribeEntity.Category = "test category"
	scribeEntity.Message = "test message"
	input <- scribeEntity
	commonEntity := <-output

	if commonEntity.ScribeEntity.Category != "test category" {
		t.Fail()
	}
	if commonEntity.ScribeEntity.Message != "test message" {
		t.Fail()
	}
}
