package config

import "strconv"

type PrismConfig struct {
	Threads int
}

func (this *PrismConfig) String() string {
	return "PrismConfig(Threads=" +
		strconv.Itoa(this.Threads) +
		")"
}

func (this *PrismConfig) SetAttributes(attributes map[string]string) {
	this.Threads, _ = strconv.Atoi(attributes["threads"])
}
