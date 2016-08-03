package main

import (
	"flag"
	"fmt"
)

var (
	customers = 2
	delay     = 0
)

func init() {
	flag.IntVar(&customers, "customers", 10, "number of customers")
	flag.IntVar(&delay, "delay", 0, "delay for each subsequent customer")
}

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
				subscriptions = append(subscriptions, request)
			default:
				for _, subscription := range subscriptions {
					subscription <- tick
				}
				tick++
			}
		}
	}()
}

// Customer doc
func Customer(delay int) {
	go func() {
		tickchan := make(SubscriptionRequest)
		TickerSubscriptionChan <- tickchan

		startTime := <-tickchan

		for i := 0; i < delay; i++ {
			startTime = <-tickchan
		}

		currentTime := startTime

		// wait for a station:
		for stationNum := 0; stationNum < 5; {
			select {
			case request := <-StationChan:
				for i := 0; i < 10; i++ {
					currentTime = <-tickchan
				}
				request <- struct{}{}
				stationNum++
			default:
				currentTime = <-tickchan
			}
		}
		stationDuration := currentTime - startTime

		// wait for a grill:
		grillStart := currentTime
		for grillNum := 0; grillNum < 1; {
			select {
			case request := <-GrillChan:
				for i := 0; i < 180; i++ {
					currentTime = <-tickchan
				}
				request <- struct{}{}
				grillNum++
			default:
				currentTime = <-tickchan
			}
		}
		grillDuration := currentTime - grillStart

		go func() {
			for {
				<-tickchan
			}
		}()

		StatChan <- Stat{currentTime - startTime, grillDuration, stationDuration}
	}()
}

// StationRequest doc
type StationRequest chan struct{}

// StationChan doc
var StationChan chan StationRequest

// Station doc
func Station() {
	go func() {
		for {
			responseChan := make(StationRequest)
			StationChan <- responseChan
			<-responseChan
		}
	}()
}

// GrillRequest doc
type GrillRequest chan struct{}

// GrillChan doc
var GrillChan chan GrillRequest

// Grill doc
func Grill() {
	go func() {
		for {
			responseChan := make(GrillRequest)
			GrillChan <- responseChan
			<-responseChan
		}
	}()
}

func main() {
	flag.Parse()

	TickerSubscriptionChan = make(chan SubscriptionRequest)
	StationChan = make(chan StationRequest)
	GrillChan = make(chan GrillRequest)
	StatChan = make(chan Stat)

	for i := 0; i < customers; i++ {
		Customer(i * delay)
	}

	for i := 0; i < 20; i++ {
		Station()
	}

	for i := 0; i < 8; i++ {
		Grill()
	}

	Ticker()

	Stats(customers)
}

// Stat doc
type Stat struct {
	TotalDuration   int
	GrillDuration   int
	StationDuration int
}

// StatChan doc
var StatChan chan Stat

// Stats doc
func Stats(customers int) {
	var durations []Stat
	for i := 0; i < customers; i++ {
		stat := <-StatChan
		durations = append(durations, stat)
	}

	var avgDuration, avgGrill, avgStation int
	for _, duration := range durations {
		avgDuration += duration.TotalDuration
		avgGrill += duration.GrillDuration
		avgStation += duration.StationDuration
	}

	fmt.Printf("Total Average Wait Time: %v\nAverage Grill Wait Time: %v\nAverage Station Wait Time: %v\n", avgDuration/customers, avgGrill/customers, avgStation/customers)
}
