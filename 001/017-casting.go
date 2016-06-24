package main

import (
    "fmt"
    "strconv"
    "reflect"
)

func main() {

    randInt := 5
    randFloat := 10.5
    randString := "100"
    randString2 := "250.5"

    fmt.Println("int to float:", float64(randInt), reflect.TypeOf(float64(randInt)))
    fmt.Println("float to int:", int(randFloat), reflect.TypeOf(int(randFloat)))

    newInt, _ := strconv.ParseInt(randString, 0, 64)
    fmt.Println("string to int:", newInt, reflect.TypeOf(newInt))

    newFloat, _ := strconv.ParseFloat(randString2, 64)
    fmt.Println("string to float:", newFloat, reflect.TypeOf(newFloat))
}
