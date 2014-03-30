package receiver

import "fmt"
import "net/http"
import "github.com/csizsek/prism/entity"
import "strconv"
import "github.com/csizsek/prism/config"

type HttpReceiver struct {
	output chan *entity.HttpEntity
	port   int
}

func (this *HttpReceiver) Receive() {
	http.HandleFunc("/", this.handler)
	http.ListenAndServe(":"+strconv.Itoa(this.port), nil)
}

func (this *HttpReceiver) handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "data received")
	entity := entity.NewHttpEntity()
	entity.Data = "baz"
	this.output <- entity
}

func (this *HttpReceiver) Configure(config config.HttpReceiverConfig) {
	this.port = config.Port
}

func NewHttpReceiver(output chan *entity.HttpEntity) *HttpReceiver {
	receiver := new(HttpReceiver)
	receiver.output = output
	return receiver
}
