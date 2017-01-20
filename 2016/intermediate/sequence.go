package main

import "fmt"

// channel for consecutive integers to assign to new connections
var idGen chan int

func init() {
	idGen = make(chan int)
	go func() {
		id := 1
		for {
			idGen <- id
			id++
		}
	}()
}

func main() {
	fmt.Println(<-idGen)
}
