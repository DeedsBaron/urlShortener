package store

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"shortener/internal/app/config"
)

var (
	Memsol string
)

type Storage interface {
	PostStore(context.Context, string, *config.Config) (string, error)
	FindInStore(context.Context, string, *config.Config) (string, error)
	Print()
}

func InitStorage(config *config.Config) Storage {
	var err error
	var newStorage Storage
	fmt.Println("type = ", config.Type)
	if config.Type == 0 {
		newStorage = NewInMemStorage()
	} else if config.Type == 1 {
		newStorage, err = NewPostgres(config)
		if err != nil {
			log.Fatal(err)
		}
		logrus.Info("Successfully connected to database")
		return newStorage
	}
	return newStorage
}
