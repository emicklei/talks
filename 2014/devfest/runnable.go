package main

import "time"

func main() {
	// START OMIT
	done := make(chan bool)

	go func() { // HL
		println("some time later")
		done <- true
	}() // HL

	select {
	case <-done:
	case <-time.After(10 * time.Second):
	}
	// END OMIT
}
