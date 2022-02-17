package main

import (
	"fmt"
	"log"
<<<<<<< HEAD
	"shortener/urlShortener/internal/app/apiserver"
=======
	"shortener/internal/app/apiserver"
>>>>>>> ed8f4a1 (postgresql container is configured and working)
)

func main() {
	config, err1 := apiserver.NewConfig()
	if err1 != nil {
		return
	}
	fmt.Println("config= ", config)
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
