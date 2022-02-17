package apiserver

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
	}
}

func NewConfig() (*Config, error) {
<<<<<<< HEAD
	flag.StringVar(&configPath, "config-path", "urlShortener/configs/apiserver.toml", "path to config file")
=======
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
>>>>>>> ed8f4a1 (postgresql container is configured and working)
	config := &Config{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}
