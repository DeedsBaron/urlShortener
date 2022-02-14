package main

import (
	"shortener/internal/app/store"
)

func main() {
	store := store.NewInMemStorage()
	store.InMemStore[1] = []string{"https", "zalupa"}
	//for ; val, ok := store.InMemStore[1] {
	//fmt.Println(rand)
	//num := uint64(randgen.Generate())
	//fmt.Println("uint64=\t\t", num)
	//str := encoder.Encode(num)
	//fmt.Println("encoded str=", str)
	//dstr, _ := encoder.Decode(str)
	//fmt.Println("decoded str=\t", dstr)
}
