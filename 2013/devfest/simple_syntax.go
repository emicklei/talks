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
	arraylist := []string{" hello", who()}

	fmt.Printf("%v on %v", arraylist, timestamp)
}
