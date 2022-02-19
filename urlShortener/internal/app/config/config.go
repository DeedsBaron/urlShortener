package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Options  struct {
		Schema string `toml:"schema"`
		Prefix string `toml:"prefix"`
	} `toml:"options"`
	Storage struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Database string `toml:"database"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Attempts int    `toml:"attempts2con"`
	} `toml:"storage"`
}

func NewConfig() (*Config, error) {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	config := &Config{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}
