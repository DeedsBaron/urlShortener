package store

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"shortener/internal/app/config"
	"shortener/internal/app/encoder"
	"shortener/internal/app/randgen"
)

type InMemStorage struct {
	InMemStore map[uint64][]string
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		InMemStore: make(map[uint64][]string),
	}
}

func (st *InMemStorage) Print() {
	fmt.Println(st.InMemStore)
}

func checkIfLongURLAlreadyInMap(m1 map[uint64][]string, longURL string) string {
	for _, URL := range m1 {
		if URL[0] == longURL {
			return URL[1]
		}
	}
	return ""
}

func validateURL(longURL string) error {
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return errors.New("Invalid URI for request")
	}
	return nil
}

func (st *InMemStorage) FindInStore(ctx context.Context, shortURL string, config *config.Config) (string, error) {
	for _, val := range st.InMemStore {
		if val[1] == config.Options.Schema+"://"+config.Options.Prefix+"/"+shortURL {
			return val[0], nil
		}
	}
	return "", errors.New("short URL doesn't exist in base")
}

func (st *InMemStorage) PostStore(ctx context.Context, longURL string, config *config.Config) (string, error) {
	if shortURL := checkIfLongURLAlreadyInMap(st.InMemStore, longURL); shortURL != "" {
		return "longURL is already in base " + shortURL, nil
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
	return st.InMemStore[rand][1], nil
}
