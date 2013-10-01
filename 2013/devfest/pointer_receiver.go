package main

import (
	"fmt"
)

// START OMIT
type Counter struct{ value int }

func (c *Counter) Inc() { c.value++ }

func increment(c Counter) { c.Inc() }

func decrement(c *Counter) { c.value-- }

func main() {
	c := Counter{}
	c.Inc()
	fmt.Println(c.value)

	increment(c)
	fmt.Println(c.value)

	decrement(&c)
	fmt.Println(c.value)
}

// END OMIT
