package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
	memSol     string
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Type     int
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

func parseFlags() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.StringVar(&memSol, "mem", "inmem", "\"inmem\" for in memory solution\n\"psql\" for postgresql solution")
	flag.Parse()
	if len(flag.Args()) != 0 {
		log.Fatal("Wrong binary parameters, try -help")
	}
}

func NewConfig() (*Config, error) {
	parseFlags()
	config := &Config{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	if memSol == "inmem" {
		config.Type = 0
	} else if memSol == "psql" {
		config.Type = 1
	} else {
		log.Fatal("Wrong memory flag, try -help")
	}
	fmt.Println(config)
	return config, nil
}
