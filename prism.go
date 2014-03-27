package main

import "github.com/csizsek/prism/entity"
import "github.com/csizsek/prism/receiver"
import "github.com/csizsek/prism/decoder"
import "github.com/csizsek/prism/dispatcher"
import "github.com/csizsek/prism/encoder"
import "github.com/csizsek/prism/sender"

func main() {
	rec2dec := make(chan entity.ScribeEntity)
	dec2disp := make(chan entity.CommonEntity)
	disp2enc := make(chan entity.CommonEntity)
	enc2send := make(chan entity.BarEntity)
	
	quit := make(chan string)
	endpoints := []chan entity.CommonEntity{disp2enc}

	receiver := receiver.NewScribeReceiver(rec2dec)
	decoder := decoder.NewScribeDecoder(rec2dec, dec2disp)
	dispatcher := dispatcher.NewDispatcher(dec2disp, endpoints)
	encoder := encoder.NewBarEncoder(disp2enc, enc2send)
	sender := sender.NewBarSender(enc2send)

	go receiver.Receive()
	go decoder.Decode()
	go dispatcher.Dispatch()
	go encoder.Encode()
	go sender.Send()

	<-quit
}
