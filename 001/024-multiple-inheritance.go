package main

import (
	"fmt"
)

func main() {
	cp := new (CameraPhone)  // a new camera phone instance
    fmt.Println("Our new CameraPhone exhibits multiple behaviors ...")
    fmt.Println("It can take a picture: ", cp.takePicture()) // exhibits behavior of a Camera
    fmt.Println("It can also make calls: ", cp.call()) // ... and also that of a Phone
}

type Camera struct {}

func (_ Camera) takePicture() string { //not using the type, so discard it by putting a _
	return "Click"
}

type Phone struct {}

func (_ Phone) call() string { // not using the type, so discard it by putting a _
	return "Ring Ring"
}

// multiple inheritance
type CameraPhone struct {
	Camera // has anonymous camera
	Phone // has anonymous phone
}
