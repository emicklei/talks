package main

import "fmt"

func main() {
	squared := make(chan int)
	go func(j int) {
		squared <- j * j
	}(256)
	fmt.Println(<-squared)
}
