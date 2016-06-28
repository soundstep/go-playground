package main

import (
    "fmt"
    "time"
)

func main() {

    rect1 := Rectangle { leftX: 0, topY: 50, width: 30, height: 10 }
    fmt.Println("rect1:", rect1);
    fmt.Println("rect1 left X:", rect1.leftX);

    rect2 := Rectangle {0, 50, 30, 10, "a name"}
    fmt.Println("rect2:", rect2);
    fmt.Println("rect2 left X:", rect2.leftX);

    fmt.Println("rect1 area:", rect2.area());

    fmt.Println("---- instantiate using new")

    pr := new (Rectangle) // get pointer to an instance with new keyword
    (*pr).width = 6 // set value using . notation by dereferencing pointer.
    pr.height = 8 // set value using . notation - same as previous.  There is no -> operator like in c++. Go automatically converts
    pr.name = "ptr_to_rectangle"
    fmt.Println("Rectangle pr as address is: ", pr) // Go performs default printing of structs
    fmt.Println("Rectangle pr as value is: ", *pr) // address and value are differentiated with an & symbol

    fmt.Println("---- anonymous fields")

    h := House{Kitchen{10}, 3} // to initialize you have to use composed type name.
    fmt.Println("House h has this many rooms:", h.numOfRooms) //numOfRooms is a field of House
    fmt.Println("House h has this many plates:", h.numOfPlates) //numOfPlates is a field of anonymous field Kitchen, so it can be referred to like a field of House
    fmt.Println("The Kitchen contents of this house are:", h.Kitchen) //we can refer to the embedded struct in its entirety by referring to the name of the struct type

    fmt.Println("---- anonymous fields conflict")

    h2 := House2{Kitchen2{2}, 10} // kitchen has 2 lamps, and the House has a total of 10 lamps
    fmt.Println("House h has this many lamps:", h2.numOfLamps) // this is ok - the outer House's numOfLamps hides the other one.  Output is 10.
    fmt.Println("The Kitchen in house h has this many lamps:", h2.Kitchen2.numOfLamps) // we can still reach the number of lamps in the kitchen by using the type name h.Kitchen

    fmt.Println("---- extend built-in")
    fmt.Println(time.Now())

    m := myTime{time.Now()}
    fmt.Println("Full time now:", m.String()) // calling existing String method on anonymous Time field
    fmt.Println("First 5 chars:", m.first10Chars()) // calling myTime.first5Chars

    fmt.Println("---- methods on anonymous")
    h3 := House3{Kitchen3{4, 4}} //the kitchen has 4 forks and 4 knives
    fmt.Println("Sum of forks and knives in house: ", h3.totalForksAndKnives())  //called on House even though the method is associated with Kitchen
}

type Rectangle struct {
    leftX float64
    topY float64
    width float64
    height float64
    name string
}

func (rect *Rectangle) area() float64 {
    return rect.width * rect.height
}

type Kitchen struct {
    numOfPlates int
}

type House struct {
    Kitchen // anonymous field
    numOfRooms int
}

type Kitchen2 struct {
    numOfLamps int
}

type House2 struct {
    Kitchen2 // anonymous field
    numOfLamps int // same field name as in the Kitchen
}

type myTime struct {
    time.Time // anonymous field
}

func (t myTime) first10Chars() string {
    return t.Time.String()[0:10]
}

type Kitchen3 struct {
    numOfForks int
    numOfKnives int
}

func (k Kitchen3) totalForksAndKnives() int {
    return k.numOfForks + k.numOfKnives
}

type House3 struct {
    Kitchen3
}
