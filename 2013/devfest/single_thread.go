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

func main() {
	ids := []int{123, 456, 789}
	offers := []string{}
	for _, each := range ids {
		offers = append(offers, fetchOffer(each)) // HL
	}
	fmt.Printf("%v\n", offers)
}
