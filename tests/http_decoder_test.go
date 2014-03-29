package main

import "testing"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/decoder"

func TestHttpDecoderDecode(t *testing.T) {
	input := make(chan *entity.HttpEntity)
	output := make(chan *entity.CommonEntity)

	httpDecoder := decoder.NewHttpDecoder(input, output)
	go httpDecoder.Decode()

	httpEntity := entity.NewHttpEntity()
	httpEntity.Data = "test data"
	input <- httpEntity
	commonEntity := <-output

	if commonEntity.HttpEntity.Data != "test data" {
		t.Fail()
	}
}
