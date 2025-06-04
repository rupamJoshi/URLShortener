package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Schema string `json:"schema"`
}

func readConfig(path string) *Config {

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}
	return config
}

func NewConfig(path string) *Config {
	return readConfig(path)
}
