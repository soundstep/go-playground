package main

import "fmt"

func main() {
	presAge := make(map[string]int)
	presAge["Roosevelt"] = 42
	fmt.Println("map:", presAge)
	fmt.Println("length:", len(presAge))

	presAge["John F, Kennedy"] = 43
	fmt.Println("map:", presAge)
	delete(presAge, "John F, Kennedy")
	fmt.Println("map:", presAge)
}
