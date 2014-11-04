package main

import (
	"fmt"
	"strconv"
)

func main() {

	// START OMIT
	/**
	 * String s = Integer.toString(42);
	 *
	 **/
	s := strconv.Itoa(42)

	// END OMIT
	fmt.Println(s)
}
