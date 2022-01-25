package main

import "fmt"

func main() {
	var x [10]int
	fmt.Println(x)
	fmt.Println("Length is", len(x))

	x[0] = 10
	x[1] = 6
	x[2] = 12
	x[3] = 16
	x[4] = 9

	fmt.Println(x)
	fmt.Println(x[1])

	z := [2]string{"Daniel", "Vinícius"}
	fmt.Println(z)
}
