package main

import "fmt"

func main() {
	yourAge := 18
	switch yourAge {
	case 16:
		fmt.Println("GO drive")
	case 18:
		fmt.Println("GO vote")
	default:
		fmt.Println("GO have fun")
	}
}
