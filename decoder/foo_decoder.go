package decoder

type FooDecoder struct {
	input chan FooEntity
	output chan CommonEntity
}

func (this *FooDecoder) decode() {
	fmt.Println("FooDecoder.decode")
	for {
		fooEntity := <- this.input
		this.output <- *NewCommonEntity(fooEntity.data, "")
	}
}

func NewFooDecoder(input chan FooEntity, output chan CommonEntity) *FooDecoder {
	fmt.Println("NewFooDecoder")
	decoder := new(FooDecoder)
	decoder.input = input
	decoder.output = output
	return decoder
}
