package main

import "fmt"
import "net/http"
import "./entity"

func main() {
	input := make(chan entity.FooEntity)
	middle := make(chan entity.CommonEntity)
	output := make(chan entity.BarEntity)
	quit := make(chan string)

	receiver := NewFooReceiver(input)
	decoder := NewFooDecoder(input, middle)
	encoder := NewBarEncoder(middle, output)
	sender := NewBarSender(output)

	go receiver.receive()
	go decoder.decode()
	go encoder.encode()
	go sender.send()

	quitMsg := <- quit
	fmt.Println(quitMsg)
}
