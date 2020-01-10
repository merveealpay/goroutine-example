package main

import (
	"fmt"
	"time"
)

func main() {
	cities := []string{"Sivas", "Erzincan", "Istanbul", "Sivas", "Izmir", "Antalya"}
	firstChannel := make(chan string)

	//finder
	go func(cities []string) {
		for _, cities := range cities {
			firstChannel <- cities //send cities to channel
			time.Sleep(time.Second)

		}

	}(cities) //expression must be function call go

	//receiver

	go func() {
		for i := 0; i < len(cities); i++ {
			found := <-firstChannel
			fmt.Println("Recipient: " + found + " received from Finder")
		}
	}() //expression must be function call go

	<-time.After(time.Second * 5) //send data to this channel after 5 seconds
}
