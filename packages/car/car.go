package car

type Car struct {
	Name                 string
	Color                string
	somePrivateAttribute string
}

func (car Car) Start() string {
	car.somePrivateMethod()
	return car.Name + " has been started"
}

func (car Car) somePrivateMethod() string {
	car.somePrivateAttribute = "secret"
	return car.somePrivateAttribute
}
