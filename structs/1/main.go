package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

// func (car Car) info() string {
// 	return fmt.Sprintf("Car: %v\n Year: %v\n Color: %v", car.Name, car.Year, car.Color)
// }

func main() {
	var cars = []Car{{"One Car", 2022, "White"}, {"Other Car", 2022, "Black"}, {"Another Car", 2022, "Red"}}

	for i := 0; i < len(cars); i++ {
		// fmt.Println(cars[i].info())
		fmt.Println(cars[i].Color)
	}
}
