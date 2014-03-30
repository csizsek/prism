package config

type StdoutSenderConfig struct {
	Prefix string
}

func (this *StdoutSenderConfig) String() string {
	return "StdoutSenderConfig(Prefix=" +
		this.Prefix +
		")"
}

func (this *StdoutSenderConfig) SetAttributes(attributes map[string]string) {
	this.Prefix = attributes["prefix"]
}
