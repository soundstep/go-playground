package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func main() {

	strbryCs := make(chan string)
	chocoCs := make(chan string)

	//two cake makers
	go makeCakeAndSend(chocoCs, "Chocolate", 3)   //make 3 chocolate cakes and send
	go makeCakeAndSend(strbryCs, "Strawberry", 3) //make 3 strawberry cakes and send

	//one cake receiver and packer
	go receiveCakeAndPack(strbryCs, chocoCs) //pack all cakes received on these cake channels

	//sleep for a while so that the program doesn’t exit immediately
	time.Sleep(4 * 1e9)
}

func makeCakeAndSend(cs chan string, flavor string, count int) {
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbryCs chan string, chocoCs chan string) {
	strbryClosed, chocoClosed := false, false

	for {
		//if both channels are closed then we can stop
		if strbryClosed && chocoClosed {
			return
		}
		fmt.Println("Waiting for a new cake ...")
		select {
		case cakeName, strbryOk := <-strbryCs:
			if !strbryOk {
				strbryClosed = true
				fmt.Println(" ... Strawberry channel closed!")
			} else {
				fmt.Println("Received from Strawberry channel.  Now packing", cakeName)
			}
		case cakeName, chocoOk := <-chocoCs:
			if !chocoOk {
				chocoClosed = true
				fmt.Println(" ... Chocolate channel closed!")
			} else {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
			}
		}
	}
}
