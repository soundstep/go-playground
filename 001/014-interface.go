package main

import (
	"fmt"
	"math"
)

func main() {

	rect := Rectangle{20, 50}
	circ := Circle{4}
	squa := Square{8}

	fmt.Println("Rectangle area:", getArea(rect))
	fmt.Println("Circle area:", getArea(circ))
	fmt.Println("Square area:", getArea(squa))

	fmt.Println("----- Looping through shapes for area ...")
	shapesArr := [...]Shape{rect, circ, squa}
	for n := range shapesArr {
		fmt.Println("Shape details: ", shapesArr[n])
		fmt.Println("Area of this shape is: ", shapesArr[n].area())
	}
}

// Shape interface
type Shape interface {
	area() float64
}

// Rectangle struct
type Rectangle struct {
	width  float64
	height float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Square struct
type Square struct {
	side float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s Square) area() float64 {
	return s.side * s.side
}

func getArea(shape Shape) float64 {
	return shape.area()
}
