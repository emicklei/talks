package main

import (
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// START OMIT
type Converters struct {
	IntegerToString func(i int) string
	StringToInteger func(s string) int
}

func main() {
	c := Converters{
		IntegerToString: strconv.Itoa,
		StringToInteger: func(s string) int {
			if i, err := strconv.Atoi(s); err != nil {
				return 0
			} else {
				return i
			}
		},
	}
	spew.Dump(c)
	fmt.Println(c.StringToInteger("42"))
}

// END OMIT
