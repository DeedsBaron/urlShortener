package store

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"shortener/urlShortener/internal/app/encoder"
	"shortener/urlShortener/internal/app/randgen"
)

var (
	Memsol string
)

type Storage interface {
	PostStore(string, string, string) (string, error)
	FindInStore(string, string, string) (string, error)
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

func (st *InMemStorage) FindInStore(shortURL, schema, prefix string) (string, error) {
	for _, val := range st.InMemStore {
		if val[1] == schema+"://"+prefix+"/"+shortURL {
			return val[0], nil
		}
	}
	return "", errors.New("Short URL doesn't exist")
}

func (st *InMemStorage) PostStore(s string, schema string, prefix string) (string, error) {
	if checkIfLongURLAlreadyInMap(st.InMemStore, s) == true {
		return "", errors.New("URL is already in base")
	}
	if err := validateURL(s); err != nil {
		return "", err
	}
	var rand uint64
	for {
		rand = randgen.Generate()
		if _, ok := st.InMemStore[rand]; ok == false {
			st.InMemStore[rand] = []string{s, schema + "://" + prefix + "/" + encoder.Encode(rand)}
			break
		}
	}
	fmt.Println("map = ", st.InMemStore)
	return st.InMemStore[rand][1], nil
}

func InitStorage() Storage {
	flag.StringVar(&Memsol, "mem", "inmem", "\"inmem\" for in memory solution\n\"psql\" for postgresql solution")
	flag.Parse()
	var newStorage Storage
	if Memsol == "inmem" {
		fmt.Println("inmem sol")
		newStorage = NewInMemStorage()
	} else if Memsol == "psql" {
		fmt.Println("psql sol")
		return nil
	} else {
		log.Fatal("Wrong mem flag")
	}
	return newStorage
}
