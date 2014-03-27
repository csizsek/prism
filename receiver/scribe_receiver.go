package receiver

import "fmt"
import "net"
import "net/rpc"
import "github.com/csizsek/prism/entity"
import "github.com/prezi/go-thrift/examples/scribe"
import "github.com/prezi/go-thrift/thrift"
import "github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"

type ScribeReceiver struct {
	output     chan entity.ScribeEntity
	lineRegExp pcre.Regexp
	port       int
	listener   net.Listener
}

func (this *ScribeReceiver) Receive() {
	for {
		conn, _ := this.listener.Accept()
		go rpc.ServeCodec(thrift.NewServerCodec(thrift.NewFramedReadWriteCloser(conn, 100000000), thrift.NewBinaryProtocol(true, false)))
	}
}

func NewScribeReceiver(output chan entity.ScribeEntity) *ScribeReceiver {
	receiver := new(ScribeReceiver)
	receiver.output = output
	receiver.lineRegExp = pcre.MustCompile(`^(?P<TimeStamp>\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d,\d{3}) (\S+) (?P<Category>\S+) (?P<Message>.+)$`, 0)
	rpc.RegisterName("Thrift", &scribe.ScribeServer{receiver})
	receiver.port = 1234
	rpc.RegisterName("Thrift", &scribe.ScribeServer{receiver})
	receiver.listener, _ = net.Listen("tcp", "127.0.0.1:1234")
	return receiver
}

func (this *ScribeReceiver) Log(messages []*scribe.LogEntry) (scribe.ResultCode, error) {
	for _, msg := range messages {
		scribeEntity := this.parseMessage(msg)
		this.output <- *scribeEntity
	}
	return scribe.ResultCodeOk, nil
}

func (this *ScribeReceiver) parseMessage(msg *scribe.LogEntry) *entity.ScribeEntity {
	scribeEntity := entity.NewScribeEntity()
	if m := this.lineRegExp.MatcherString(msg.Message, 0); m != nil {
		scribeEntity.Category = m.GroupString(3)
		scribeEntity.Message = m.GroupString(4)
	} else {
		fmt.Println("no match :(")
		fmt.Println(msg.Message)
	}
	return scribeEntity
}
