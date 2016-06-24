package main

import "fmt"

func main() {
    //------------- defer

    defer printTwo()
    printOne()

    //------------- safe division

    fmt.Println(safeDiv(3, 0));
    fmt.Println(safeDiv(3, 2));
}

func printOne() {
    fmt.Println("print 1")
}

func printTwo() {
    fmt.Println("print 2")
}

func safeDiv(num1, num2 int) int {
    defer func() {
        fmt.Println(recover())
    }()
    solution := num1 / num2
    return solution
}
