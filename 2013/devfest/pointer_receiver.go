package main

import (
	"fmt"
)

// START OMIT
type Counter struct{ value int }

func (c *Counter) inc() { c.value++ }

func increment(c Counter) { c.inc() }

func decrement(c *Counter) { c.value-- }

func main() {
	c := Counter{0}
	c.inc()
	fmt.Println(c.value)

	increment(c)
	fmt.Println(c.value)

	decrement(&c)
	fmt.Println(c.value)
}

// END OMIT
