package main

import "testing"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/decoder"

func TestFooDecoderDecode(t *testing.T) {
	input := make(chan *entity.FooEntity)
	output := make(chan *entity.CommonEntity)

	fooDecoder := decoder.NewFooDecoder(input, output)
	go fooDecoder.Decode()

	fooEntity := entity.NewFooEntity()
	fooEntity.Data = "test data"
	input <- fooEntity
	commonEntity := <-output

	if commonEntity.FooEntity.Data != "test data" {
		t.Fail()
	}
}
