package encoder

type Encoder interface {
	encode(CommonEntity) Entity
}
