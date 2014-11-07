package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()

	p(now)
	p(now.Add(1 * time.Hour))
	p(now.Weekday())

	then, _ := time.Parse(time.RFC3339, "2003-09-16T15:04:05Z")

	p(int(now.Sub(then).Hours()/float64(365*24)), " years")
}
