package main

import "fmt"

func main() {
    // without pointer
    x := 0
    changeXVal(x)
    fmt.Println("x:", x)
    // with pointer
    changeXValNow(&x)
    fmt.Println("x:", x)
    fmt.Println("Memory address for x:", &x)
    // send pointer directly
    yPointer := new(int)
    changeYValNow(yPointer)
    fmt.Println("y:", *yPointer)
}

func changeXVal(x int) {
    x = 2
}

func changeXValNow(x *int) {
    *x = 2
}

func changeYValNow(yPointer *int) {
    *yPointer = 100
}
