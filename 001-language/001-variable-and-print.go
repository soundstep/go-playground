package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var age = 40
	var favNum = 1.6180339
	fmt.Println(age, favNum)
	fmt.Printf("%d \n", age)

	var numOne = 1.000
	var num99 = .999
	fmt.Println(numOne - num99)

	const pi float64 = 3.14159265
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", pi)
	fmt.Printf("%e \n", pi)

	var (
		varA = 2
		varB = 3
	)
	fmt.Println(varA, varB)

	var myName = "a very long name"
	fmt.Println(len(myName))

	var isOver40 = true
	fmt.Println(isOver40)
	fmt.Printf("%t \n", isOver40)

}
