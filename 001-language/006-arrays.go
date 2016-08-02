package main

import "fmt"

func main() {
	var favNums2 [5]float64
	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618
	fmt.Println(favNums2)

	favNums3 := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(favNums3)

	for i, value := range favNums3 {
		fmt.Println(i, value)
	}
	for _, val := range favNums3 {
		fmt.Println(val)
	}
}
