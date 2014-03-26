package receiver

type FooReceiver struct {
	output chan FooEntity
}

func (this *FooReceiver) receive() {
	fmt.Println("FooReceiver.receive")
	http.HandleFunc("/", this.handler)
	http.ListenAndServe(":8000", nil)
}

func (this *FooReceiver) handler(rw http.ResponseWriter,req *http.Request) {
	fmt.Println("FooReceiver.handler")
	fmt.Fprintf(rw, "ACK")
	this.output <- *NewFooEntity("hello")
}

func NewFooReceiver(output chan FooEntity) *FooReceiver {
	fmt.Println("NewFooReceiver")
	receiver := new(FooReceiver)
	receiver.output = output
	return receiver
}
