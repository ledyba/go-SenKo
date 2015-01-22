package main

import (
	"fmt"
	"senko"
)

func main() {
	mat := senko.LoadMatrix("data/matrix.def")
	dict := senko.NewDictionary()
	dict.LoadDictionaries("data")
	dec := senko.NewDecoder(dict, mat)
	str := "我輩は狐である"
	fmt.Printf("Decoding: %s\n", str);
	for _, s := range dec.Search(str) {
		fmt.Printf("%s\n", s.Repr())
	}
}
