package main

import "fmt"

func main() {
	yourAge := 18
	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can have fun")
	}
}
