package main

import "fmt"

type vehicle interface {
	start() string
}

func startVehicle(vehicle vehicle) string {
	return vehicle.start()
}

type car struct {
	name string
}

type motorcycle struct {
	name string
}

func (car car) start() string {
	return "The car " + car.name + " has been started"
}

func (motorcycle motorcycle) start() string {
	return "The motorcycle " + motorcycle.name + " has been started"
}

func main() {
	car := car{"Gol"}
	motorcycle := motorcycle{"Honda"}

	fmt.Println(startVehicle(car))
	fmt.Println(startVehicle(motorcycle))
}
