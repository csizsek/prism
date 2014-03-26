package main

import "fmt"

import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/receiver"
import "github.com/csizsek/prism/decoder"
import "github.com/csizsek/prism/encoder"
import "github.com/csizsek/prism/sender"

func main() {
	input := make(chan entity.FooEntity)
	middle := make(chan entity.CommonEntity)
	output := make(chan entity.BarEntity)
	quit := make(chan string)

	receiver := receiver.NewFooReceiver(input)
	decoder := decoder.NewFooDecoder(input, middle)
	encoder := encoder.NewBarEncoder(middle, output)
	sender := sender.NewBarSender(output)

	go receiver.Receive()
	go decoder.Decode()
	go encoder.Encode()
	go sender.Send()

	quitMsg := <-quit
	fmt.Println(quitMsg)
}
