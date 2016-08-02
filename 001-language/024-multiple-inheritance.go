package main

import (
	"fmt"
)

func main() {
	cp := new(CameraPhone) // a new camera phone instance
	fmt.Println("Our new CameraPhone exhibits multiple behaviors ...")
	fmt.Println("It can take a picture: ", cp.takePicture()) // exhibits behavior of a Camera
	fmt.Println("It can also make calls: ", cp.call())       // ... and also that of a Phone
}

// Camera struct
type Camera struct{}

func (c Camera) takePicture() string { //not using the type, so discard it by putting a _
	return "Click"
}

// Phone struct
type Phone struct{}

func (p Phone) call() string { // not using the type, so discard it by putting a _
	return "Ring Ring"
}

// CameraPhone multiple inheritance
type CameraPhone struct {
	Camera // has anonymous camera
	Phone  // has anonymous phone
}
