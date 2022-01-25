package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Color string `json:"color"`
}

func main() {
	var car Car
	data := []byte(`{"name": "Gol", "year": 2022, "color": "Yellow"}`)

	// fmt.Println(data, string(data))

	json.Unmarshal(data, &car)
	fmt.Println(car.Name, car.Year, car.Color)
}
