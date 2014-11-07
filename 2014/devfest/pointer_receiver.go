package main

import (
	"fmt"
)

// START OMIT
type Counter struct {
	value int
}

func (c *Counter) Inc() {
	c.value++
}
func (c Counter) IncNoEffect() {
	c.value++
}
func main() {
	c := Counter{0}
	c.Inc()
	fmt.Println(c.value)

	c.IncNoEffect()
	fmt.Println(c.value)
}

// END OMIT
