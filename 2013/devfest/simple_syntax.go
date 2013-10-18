package main

import (
	"fmt"
	"time"
)

func who() string {
	return "bol.com ™ "
}

func main() {
	timestamp := time.Now()
	slice := []string{" hello", who()} // HL

	fmt.Printf("%v on %v", slice, timestamp)
}
