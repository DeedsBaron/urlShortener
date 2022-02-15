package main

import (
	"fmt"
	"shortener/internal/app/encoder"
	"shortener/internal/app/randgen"
)

func main() {
	//for ; val, ok := store.InMemStore[1] {
	rand := randgen.Generate()
	fmt.Println(rand)
	num := uint64(randgen.Generate())
	fmt.Println("uint64=\t\t", num)
	str := encoder.Encode(num)
	fmt.Println("encoded str=", str)
	dstr, _ := encoder.Decode(str)
	fmt.Println("decoded str=\t", dstr)
}
