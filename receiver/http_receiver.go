package receiver

import "fmt"
import "net/http"
import "github.com/csizsek/prism/entity"

type HttpReceiver struct {
	output chan *entity.HttpEntity
}

func (this *HttpReceiver) Receive() {
	http.HandleFunc("/", this.handler)
	http.ListenAndServe(":8000", nil)
}

func (this *HttpReceiver) handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "data received")
	entity := entity.NewHttpEntity()
	entity.Data = "baz"
	this.output <- entity
}

func NewHttpReceiver(output chan *entity.HttpEntity) *HttpReceiver {
	receiver := new(HttpReceiver)
	receiver.output = output
	return receiver
}
