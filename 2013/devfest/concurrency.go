package main

import (
	"fmt"
	"time"
)

func fetchOffer(id int) string {
	fmt.Printf("fetching offer for id:%d\n", id)
	time.Sleep(2 * time.Second)
	return "<offer/>"
}

var ids = []int{123, 456, 789}
var offers = []string{}

// START OMIT
func fetch(id int, output chan string) {
	output <- fetchOffer(id) // HL
}

func main() {
	offer_channel := make(chan string, 10) // HL

	for _, each := range ids {
		go fetch(each, offer_channel) // HL
	}

	for i := 0; i < len(ids); i++ {
		offer := <-offer_channel // HL
		offers = append(offers, offer)
	}
	// END OMIT

	fmt.Printf("%v\n", offers)
}
