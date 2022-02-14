package store

import (
	"flag"
	"fmt"
	"log"
	"shortener/internal/app/encoder"
	"shortener/internal/app/randgen"
)

var (
	Memsol string
)

type Storage interface {
	PostStore(string, string, string) string
}

type InMemStorage struct {
	InMemStore map[uint64][]string
}

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		InMemStore: make(map[uint64][]string),
	}
}

func (st *InMemStorage) PostStore(s string, schema string, prefix string) string {
	rand := randgen.Generate()
	for {
		if _, ok := st.InMemStore[rand]; ok == false {
			st.InMemStore[rand] = []string{s, schema + "://" + prefix + "/" + encoder.Encode(rand)}
			break
		}
	}

	fmt.Println("map[1] = ", st.InMemStore[rand][1])
	fmt.Println("map = ", st.InMemStore)

	return st.InMemStore[rand][1]
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
