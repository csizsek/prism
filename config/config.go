package config

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type FrontendConfig struct {
	Protocol string
	HttpReceiverConfig
}

func (this *FrontendConfig) String() string {
	return "FrontendConfig(" +
		"Protocol=" + this.Protocol +
		", HttpReceiverConfig=" + this.HttpReceiverConfig.String() +
		")"
}

type BackendConfig struct {
	Protocol string
	StdoutSenderConfig
}

func (this *BackendConfig) String() string {
	return "BackendConfig(" +
		"Protocol=" + this.Protocol +
		", StdoutReceiverConfig=" + this.StdoutSenderConfig.String() +
		")"
}

type Config struct {
	PrismConfig
	FrontendConfig
	Backends []BackendConfig
}

func (this *Config) String() string {
	ret := "Config(" +
		"PrismConfig=" + this.PrismConfig.String() +
		", FrontendConfig=" + this.FrontendConfig.String() +
		", Backends=["
	backends := ""
	for i, cfg := range this.Backends {
		if i > 0 {
			backends += ", "
		}
		backends += cfg.String()
	}
	ret += backends
	ret += "])"
	return ret
}

func ParseConfig() *Config {
	file, err := os.Open("prism.cfg")
	if err != nil {
		println("111!1!!!!")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sectionRegexp := regexp.MustCompile(`^\[(.*)\]$`)
	section := ""
	sectionAttributeRegexp := regexp.MustCompile(`^(.*)\=(.*)$`)
	sectionAttributes := make(map[string]string)
	sectionProtocolRegexp := regexp.MustCompile(`^protocol=(.*)$`)
	sectionProtocol := ""
	var config Config
	config.Backends = make([]BackendConfig, 0, 5)
	for scanner.Scan() {
		line := strings.Trim(string(scanner.Bytes()), "\n\t ")
		match := sectionRegexp.MatchString(line)
		if match {
			if section != "" {
				updateConfig(&config, section, sectionProtocol, sectionAttributes)
			}
			section = sectionRegexp.FindStringSubmatch(line)[1]
			sectionProtocol = ""
			sectionAttributes = make(map[string]string)
			continue
		}
		match = sectionProtocolRegexp.MatchString(line)
		if match {
			sectionProtocol = sectionProtocolRegexp.FindStringSubmatch(line)[1]
			continue
		}
		match = sectionAttributeRegexp.MatchString(line)
		if match {
			key := sectionAttributeRegexp.FindStringSubmatch(line)[1]
			value := sectionAttributeRegexp.FindStringSubmatch(line)[2]
			sectionAttributes[key] = value
		}
	}
	updateConfig(&config, section, sectionProtocol, sectionAttributes)
	println(config.String())
	return &config
}

func updateConfig(config *Config, section string, sectionProtocol string, sectionAttributes map[string]string) {
	switch section {
	case "prism":
		var prismConfig PrismConfig
		prismConfig.SetAttributes(sectionAttributes)
		config.PrismConfig = prismConfig
	case "receiver":
		switch sectionProtocol {
		case "http":
			var httpReceiverConfig HttpReceiverConfig
			httpReceiverConfig.SetAttributes(sectionAttributes)
			var frontendConfig FrontendConfig
			frontendConfig.Protocol = sectionProtocol
			frontendConfig.HttpReceiverConfig = httpReceiverConfig
			config.FrontendConfig = frontendConfig
		default:
			println("unknown receiver protocol")
		}
	case "sender":
		switch sectionProtocol {
		case "stdout":
			var stdoutSenderConfig StdoutSenderConfig
			stdoutSenderConfig.SetAttributes(sectionAttributes)
			var backendConfig BackendConfig
			backendConfig.Protocol = sectionProtocol
			backendConfig.StdoutSenderConfig = stdoutSenderConfig
			config.Backends = append(config.Backends, backendConfig)
		default:
			println("unknown sender protocol")
		}
	}
}
