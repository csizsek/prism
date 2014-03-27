package receiver

import "fmt"
import "net/http"
import "github.com/csizsek/prism/entity"

type FooReceiver struct {
	output chan *entity.FooEntity
}

func (this *FooReceiver) Receive() {
	http.HandleFunc("/", this.handler)
	http.ListenAndServe(":8000", nil)
}

func (this *FooReceiver) handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "data received")
	entity := entity.NewFooEntity()
	entity.Data = "baz"
	this.output <- entity
}

func NewFooReceiver(output chan *entity.FooEntity) *FooReceiver {
	receiver := new(FooReceiver)
	receiver.output = output
	return receiver
}
