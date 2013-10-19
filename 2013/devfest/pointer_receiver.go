package main

import (
	"fmt"
)

// START OMIT
type Counter struct {
	value int
}

func increment(c *Counter) {
	c.value++
}

func main() {
	c := Counter{0}
	fmt.Println(c.value)

	increment(&c)
	fmt.Println(c.value)
}

// END OMIT
