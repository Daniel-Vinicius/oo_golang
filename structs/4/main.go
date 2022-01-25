package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"model"`
	Year  int    `json:"-"`
	Color string `json:"color"`
}

func main() {
	car := Car{"Gol", 2022, "Yellow"}

	result, _ := json.Marshal(car)
	carConvertedToJSON := string(result)

	fmt.Println("Car on JSON", carConvertedToJSON)
}
