package main

import "fmt"

type Names []interface{}

func (names *Names) load() {
	*names = Names{
		"Daniel",
		"Wesley",
		"Davi",
		1,
	}
}

func (names Names) printNames() {
	fmt.Println(names)
}

func main() {
	var names Names
	names.load()
	names.printNames()
}
