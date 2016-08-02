package main

import (
	"fmt"
)

func main() {
	f := Ferrari{Car{4}}
	fmt.Println("A Ferrari has this many wheels: ", f.numberOfWheels()) // has car behavior
	f.sayHiToSchumacher()                                               // has Ferrari behavior

	a := AstonMartin{Car{4}}
	fmt.Println("An Aston Martin has this many wheels: ", a.numberOfWheels()) // has car behavior
	a.sayHiToBond()                                                           // has AstonMartin behavior
}

// Car struct
type Car struct {
	wheelCount int
}

// define a behavior for Car
func (car Car) numberOfWheels() int {
	return car.wheelCount
}

// Ferrari struct
type Ferrari struct {
	Car // anonymous  field Car
}

// a behavior only available for the Ferrari
func (f Ferrari) sayHiToSchumacher() {
	fmt.Println("Hi Schumacher!")
}

// AstonMartin struct
type AstonMartin struct {
	Car
}

// a behavior only available for the AstonMartin
func (a AstonMartin) sayHiToBond() {
	fmt.Println("Hi Bond, James Bond!")
}
