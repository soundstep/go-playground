package main

import "fmt"

func main() {
	numSlice := []int{5, 4, 3, 2, 1}
	fmt.Println("numSlice[3:5]:", numSlice[3:5])
	fmt.Println("numSlice[:2]:", numSlice[:2])
	fmt.Println("numSlice[2:]:", numSlice[2:])

	numSlice3 := make([]int, 5, 10)
	copy(numSlice3, numSlice)
	fmt.Println("copy:", numSlice3)

	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println("append:", numSlice3)
}
