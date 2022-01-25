package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Programmer struct {
	Person
	ProgrammingLanguages []string
}

func (person Person) IslegalAge() bool {
	if person.Age >= 18 {
		return true
	} else {
		return false
	}
}

func main() {
	daniel := Person{"Daniel Vinícius", 14}
	programmingLanguages := []string{"Golang", "Javascript"}

	programmer := Programmer{
		Person:               daniel,
		ProgrammingLanguages: programmingLanguages,
	}

	fmt.Println(daniel)
	fmt.Println(programmer)

	fmt.Println(daniel.IslegalAge())
	fmt.Println(programmer.IslegalAge())

	daniel.Age = 19
	fmt.Println(daniel.IslegalAge())
	fmt.Println(programmer.IslegalAge())

	programmer.Age = 19
	fmt.Println(programmer.IslegalAge())
}
