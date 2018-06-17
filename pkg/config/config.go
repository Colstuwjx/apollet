package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	User    string
	Pid     string
	Debug   bool
	Env     string
	Cluster string
}

func NewConf(confPath string) (*Config, error) {
	return local(confPath)
}

func local(confPath string) (*Config, error) {
	config := &Config{}
	_, err := toml.DecodeFile(confPath, config)

	return config, err
}
