package encoder

type BarEncoder struct {
	input chan CommonEntity 
	output chan BarEntity
}

func (this *BarEncoder) encode() {
	fmt.Println("BarEncoder.encode")
	for {
		commonEntity := <- this.input
		this.output <- *NewBarEntity(commonEntity.foo_data)
	}
}

func NewBarEncoder(input chan CommonEntity, output chan BarEntity) *BarEncoder {
	fmt.Println("NewBarEncoder")
	encoder := new(BarEncoder)
	encoder.input = input
	encoder.output = output
	return encoder
}
