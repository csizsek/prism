package encoder

import "github.com/csizsek/prism/entity"

type Encoder interface {
	Encode(*entity.CommonEntity) *entity.Entity
}
