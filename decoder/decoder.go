package decoder

type Decoder interface {
	decode(Entity) CommonEntity
}
