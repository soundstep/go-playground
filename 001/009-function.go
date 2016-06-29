package main

import "fmt"

func main() {
	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum:", addThemUp(listNums))

	num1, num2 := next2Values(5)
	fmt.Println("Next 2 values:", num1, num2)

	fmt.Println("Substract:", substractThem(1, 2, 3, 4, 5))

	//------------- closure

	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		return num3
	}
	fmt.Println("Closure:", doubleNum())
	fmt.Println("Closure:", doubleNum())

	//------------- recursion

	fmt.Println("Recursion:", factorial(3))

	//------------- naked return

	fmt.Println(split(17))

}

func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func next2Values(number int) (int, int) {
	return number + 1, number + 2
}

func substractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
