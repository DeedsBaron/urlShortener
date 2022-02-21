package main

import (
	"log"
	"shortener/internal/app/apiserver"
	"shortener/internal/app/config"
)

func main() {
	config, err1 := config.NewConfig()
	if err1 != nil {
		return
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
