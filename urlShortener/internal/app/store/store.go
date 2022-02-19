package store

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/url"
	"shortener/internal/app/config"
	"shortener/internal/app/encoder"
	"shortener/internal/app/randgen"
)

var (
	Memsol string
)

type Storage interface {
	PostStore(context.Context, string, *config.Config) (string, error)
	FindInStore(context.Context, string, *config.Config) (string, error)
}

type InMemStorage struct {
	InMemStore map[uint64][]string
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		InMemStore: make(map[uint64][]string),
	}
}

func checkIfLongURLAlreadyInMap(m1 map[uint64][]string, longURL string) bool {
	for _, URL := range m1 {
		fmt.Println("URL = ", URL[0])
		if URL[0] == longURL {
			return true
		}
	}
	return false
}

func validateURL(longURL string) error {
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return err
	}
	return nil
}

func (st *InMemStorage) FindInStore(ctx context.Context, shortURL string, config *config.Config) (string, error) {
	for _, val := range st.InMemStore {
		if val[1] == config.Options.Schema+"://"+config.Options.Prefix+"/"+shortURL {
			return val[0], nil
		}
	}
	return "", errors.New("short URL doesn't exist")
}

func (st *InMemStorage) PostStore(ctx context.Context, longURL string, config *config.Config) (string, error) {
	if checkIfLongURLAlreadyInMap(st.InMemStore, longURL) == true {
		return "", errors.New("URL is already in base")
	}
	if err := validateURL(longURL); err != nil {
		return "", err
	}
	var rand uint64
	for {
		rand = randgen.Generate()
		if _, ok := st.InMemStore[rand]; ok == false {
			st.InMemStore[rand] = []string{longURL, config.Options.Schema + "://" + config.Options.Prefix + "/" + encoder.Encode(rand)}
			break
		}
	}
	fmt.Println("map = ", st.InMemStore)
	return st.InMemStore[rand][1], nil
}

func InitStorage(config *config.Config) Storage {
	flag.StringVar(&Memsol, "mem", "inmem", "\"inmem\" for in memory solution\n\"psql\" for postgresql solution")
	flag.Parse()
	var err error
	var newStorage Storage
	if Memsol == "inmem" {
		fmt.Println("inmem sol")
		newStorage = NewInMemStorage()
	} else if Memsol == "psql" {
		newStorage, err = NewPostgres(config)
		if err != nil {
			log.Fatal(err)
		}
		logrus.Info("Successfully connected to database")
		return newStorage
	} else {
		log.Fatal("Wrong mem flag")
	}
	return newStorage
}
