package main

import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/receiver"
import "github.com/csizsek/prism/decoder"
import "github.com/csizsek/prism/encoder"
import "github.com/csizsek/prism/sender"

func main() {
	input := make(chan entity.ScribeEntity)
	middle := make(chan entity.CommonEntity)
	output := make(chan entity.BarEntity)
	quit := make(chan string)

	receiver := receiver.NewScribeReceiver(input)
	decoder := decoder.NewScribeDecoder(input, middle)
	encoder := encoder.NewBarEncoder(middle, output)
	sender := sender.NewBarSender(output)

	go receiver.Receive()
	go decoder.Decode()
	go encoder.Encode()
	go sender.Send()

	<-quit
}
