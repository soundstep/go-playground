package main

import "fmt"

func main() {

    rect1 := Rectangle { leftX: 0, topY: 50, width: 30, height: 10 }
    fmt.Println("rect1:", rect1);
    fmt.Println("rect1 left X:", rect1.leftX);

    rect2 := Rectangle {0, 50, 30, 10}
    fmt.Println("rect2:", rect2);
    fmt.Println("rect2 left X:", rect2.leftX);

    fmt.Println("rect1 area:", rect2.area());

}

type Rectangle struct {
    leftX float64
    topY float64
    width float64
    height float64
}

func (rect *Rectangle) area() float64 {
    return rect.width * rect.height
}
