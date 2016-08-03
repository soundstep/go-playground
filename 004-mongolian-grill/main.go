package main

import (
	"fmt"
	"time"
)

// SubscriptionRequest doc
type SubscriptionRequest chan int

// TickerSubscriptionChan doc
var TickerSubscriptionChan chan SubscriptionRequest

// Ticker doc
func Ticker() {
	go func() {
		tick := 0
		var subscriptions []SubscriptionRequest
		for {
			select {
			case request := <-TickerSubscriptionChan:
				fmt.Println("Append ticker subscription", request)
				subscriptions = append(subscriptions, request)
			default:
				for _, subscription := range subscriptions {
					fmt.Println("Update ticker subscription", tick, subscription)
					subscription <- tick
				}
				tick++
			}
		}
	}()
}

// Customer doc
func Customer() {
	go func() {
		fmt.Println("Customer created!")
		tickchan := make(SubscriptionRequest)
		TickerSubscriptionChan <- tickchan

		startTime := <-tickchan
		currentTime := startTime

		// wait for a station
		for stationNum := 0; stationNum < 5; {
			fmt.Println("station num:", stationNum)
			select {
			case request := <-StationChan:
				for i := 0; i < 10; i++ {
					currentTime = <-tickchan
					fmt.Println("station num:", stationNum, "current time:", currentTime)
				}
				request <- struct{}{}
				stationNum++
			default:
				currentTime = <-tickchan
			}
		}

		fmt.Printf("50 == %v\n", currentTime-startTime)
	}()
}

// StationRequest doc
type StationRequest chan struct{}

// StationChan doc
var StationChan chan StationRequest

// Station doc
func Station() {
	go func() {
		responseChan := make(StationRequest)
		StationChan <- responseChan
		<-responseChan
	}()
}

func testchan() {
	go func() {
		fmt.Println("TEST CHAN")
		tickchan := make(SubscriptionRequest)
		TickerSubscriptionChan <- tickchan
		tick1 := <-tickchan
		fmt.Println("tick1", tick1)
		// tickchan := make(SubscriptionRequest)
		// TickerSubscriptionChan <- tickchan
		// tick1 := <-tickchan
		// fmt.Println("tick1", tick1)
		// TickerSubscriptionChan <- tickchan
	}()
}

func main() {
	fmt.Println("----- Mongolian Grill -----")
	TickerSubscriptionChan = make(chan SubscriptionRequest)
	StationChan = make(chan StationRequest)

	// Customer()
	//
	// for i := 0; i < 20; i++ {
	// 	Station()
	// }
	//
	// Ticker()

	// testchan()

	//
	// startTime := <-tickchan
	// fmt.Println("start time:", startTime)

	time.Sleep(2 * time.Second)
}
