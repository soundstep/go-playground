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
    fmt.Println()

    var i int
    fmt.Println("value of i is: ", i)
    fmt.Println("address of i is: ", &i)
    fmt.Println("value at address ", &i, " is: ", *(&i)) //value at (address of i)
    fmt.Println()
    var s string
    fmt.Println("value of s is: ", s)
    fmt.Println("address of s is: ", &s)
    fmt.Println("value at address ", &s, " is: ", *&s) ////value at address of i
    fmt.Println()
    var f float64
    fmt.Println("value of f is: ", f)
    fmt.Println("address of f is: ", &f)
    fmt.Println("value at address ", &f, " is: ", *&f)
    fmt.Println()
    var c complex64
    fmt.Println("value of c is: ", c)
    ptr := &c //address of c.
    fmt.Println("address of c is: ", ptr)
    fmt.Println("value at address ", ptr, " is: ", *ptr) //value at the address
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
