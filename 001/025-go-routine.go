package main

import (
	"fmt"
	"time"
)

func main() {
	go simulateEvent("100m sprint", 10) // start 100m sprint, it should take 10 seconds
	go simulateEvent("Long jump", 6)    // start long jump, it should take 6 seconds
	go simulateEvent("High jump", 3)    // start high jump, it should take 3 seconds
	time.Sleep(time.Duration(12) * time.Second)
}

func simulateEvent(name string, timeInSecs int64) {
	// sleep for a while to simulate time consumed by event
	fmt.Println("Started ", name, ": Should take", timeInSecs, "seconds.")
	time.Sleep(time.Duration(timeInSecs) * time.Second)
	fmt.Println("Finished ", name)
}
