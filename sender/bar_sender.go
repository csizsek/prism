package sender

type BarSender struct {
	input chan BarEntity
}

func (this *BarSender) send() {
	fmt.Println("BarSender.send")
	var e BarEntity
	for {
		e = <- this.input
		fmt.Println(e.msg)
	}
}

func NewBarSender(input chan BarEntity) *BarSender {
	fmt.Println("NewBarSender")
	sender := new(BarSender)
	sender.input = input
	return sender
}
