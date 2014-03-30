package config

import "strconv"

type HttpReceiverConfig struct {
	Port int
}

func (this *HttpReceiverConfig) String() string {
	return "HttpReceiverConfig(Port=" +
		strconv.Itoa(this.Port) +
		")"
}

func (this *HttpReceiverConfig) SetAttributes(attributes map[string]string) {
	this.Port, _ = strconv.Atoi(attributes["port"])
}
