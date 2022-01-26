package main

import (
	"fmt"

	"github.com/Daniel-Vinicius/oo_golang/packages/car"
)

func main() {
	myCar := car.Car{
		Name:  "Gol",
		Color: "Black",
	}

	/*
		myCar.somePrivateAttribute
		myCar.somePrivateMethod
		The code above gives error, somePrivateAttribute and somePrivateMethod are private
	*/

	fmt.Println(myCar.Start())
}
