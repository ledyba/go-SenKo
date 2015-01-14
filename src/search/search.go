package main

import (
	"fmt"
	"komari"
)

func main() {
	mat := komari.LoadMatrix("data/matrix.def")
	dict := komari.NewDictionary()
	dict.LoadDictionaries("data")
	dec := komari.NewDecoder(dict, mat)
	str := "我輩は狐である"
	fmt.Printf("Decoding: %s\n", str);
	for _, s := range dec.Search(str) {
		fmt.Printf("%s\n", s.Repr())
	}
}
