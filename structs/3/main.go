package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	color string // Lowercase private attribute
}

func main() {
	car := Car{"Gol", 2022, "Yellow"}

	/*
		Other packages cannot see car.color just this package.
		That's why carConvertedToJSON doesn't show the color, because the color is a private attribute, because it was written in lowercase
		So the json package doesn't see car.color and it doesn't show either
	*/

	fmt.Println(car.color)

	result, _ := json.Marshal(car)
	carConvertedToJSON := string(result)

	fmt.Println("Result Car ASCII Table:", result)
	fmt.Println("Car on JSON", carConvertedToJSON)
}
