package receiver

import "net"
import "net/rpc"
import "github.com/csizsek/prism/entity"
import "github.com/prezi/go-thrift/examples/scribe"
import "github.com/prezi/go-thrift/thrift"

type ScribeReceiver struct {
	output   chan *entity.ScribeEntity
	port     int
	listener net.Listener
}

func (this *ScribeReceiver) Receive() {
	for {
		conn, _ := this.listener.Accept()
		go rpc.ServeCodec(thrift.NewServerCodec(thrift.NewFramedReadWriteCloser(conn, 100000000), thrift.NewBinaryProtocol(true, false)))
	}
}

func NewScribeReceiver(output chan *entity.ScribeEntity) *ScribeReceiver {
	receiver := new(ScribeReceiver)
	receiver.output = output
	rpc.RegisterName("Thrift", &scribe.ScribeServer{receiver})
	receiver.port = 1234
	receiver.listener, _ = net.Listen("tcp", "127.0.0.1:1234")
	return receiver
}

func (this *ScribeReceiver) Log(messages []*scribe.LogEntry) (scribe.ResultCode, error) {
	for _, msg := range messages {
		scribeEntity := this.parseMessage(msg)
		this.output <- scribeEntity
	}
	return scribe.ResultCodeOk, nil
}

func (this *ScribeReceiver) parseMessage(msg *scribe.LogEntry) *entity.ScribeEntity {
	scribeEntity := entity.NewScribeEntity()
	scribeEntity.Category = msg.Category
	scribeEntity.Message = msg.Message
	return scribeEntity
}
