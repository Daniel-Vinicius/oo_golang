package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	Name   string
	CanFly bool
}

func (car Car) info() string {
	return fmt.Sprintf("Car: %v\n Year: %v\n Color: %v", car.Name, car.Year, car.Color)
}

func main() {
	fusca := Car{"Fusca", 2040, "Purple"}

	superCar := SuperCar{
		Car:    fusca,
		Name:   "Fusca of the future",
		CanFly: true,
	}

	fmt.Println(fusca.info())
	fmt.Println(superCar.info())
	fmt.Println("Fusca.info() is equal to superCar.info():", fusca.info() == superCar.info())

	fmt.Println("")

	fmt.Println("CanFly:", superCar.CanFly)
	fmt.Println("New name:", superCar.Name)
	fmt.Println("Old name:", superCar.Car.Name)
}
