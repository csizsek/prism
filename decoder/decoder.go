package decoder

import "github.com/csizsek/prism/entity"

type Decoder interface {
	Decode(entity.Entity) entity.CommonEntity
}
