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

	Http *HTTP
}

type HTTP struct {
	Host string
	Port int
}

func NewConf(confPath string) (*Config, error) {
	return local(confPath)
}

func local(confPath string) (*Config, error) {
	config := &Config{}
	_, err := toml.DecodeFile(confPath, config)

	return config, err
}
