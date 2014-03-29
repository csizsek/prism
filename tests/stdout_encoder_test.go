package main

import "testing"
import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/encoder"

func TestStdoutEncoderEncode(t *testing.T) {
	input := make(chan *entity.CommonEntity)
	output := make(chan *entity.StdoutEntity)

	stdoutEncoder := encoder.NewStdoutEncoder(input, output)
	go stdoutEncoder.Encode()

	commonEntity := entity.NewCommonEntity()
	commonEntity.StdoutEntity.Msg = "test msg"
	input <- commonEntity
	stdoutEntity := <-output

	if stdoutEntity.Msg != "test msg" {
		t.Fail()
	}
}
