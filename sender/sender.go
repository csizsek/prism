package sender

import "github.com/csizsek/prism/entity"

type Sender interface {
	Send(entity.Entity)
}
