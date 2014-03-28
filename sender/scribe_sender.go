package sender

import "net"
import "strings"
import "github.com/csizsek/prism/entity"
import "github.com/prezi/go-thrift/examples/scribe"
import "github.com/prezi/go-thrift/thrift"

type ScribeSender struct {
	input        chan *entity.ScribeEntity
	scribeClient scribe.ScribeClient
}

func (this *ScribeSender) Send() {
	for {
		scribeEntity := <-this.input
		logEntry := scribe.LogEntry{scribeEntity.Category, strings.TrimRight(scribeEntity.Message, "\n") + "\n"}
		this.scribeClient.Log([]*scribe.LogEntry{&logEntry})
	}
}

func NewScribeSender(input chan *entity.ScribeEntity) *ScribeSender {
	sender := new(ScribeSender)
	sender.input = input
	conn, _ := net.Dial("tcp", "127.0.0.1:1463")
	sender.scribeClient.Client = thrift.NewClient(thrift.NewFramedReadWriteCloser(conn, 0), thrift.NewBinaryProtocol(true, false), false)

	return sender
}
