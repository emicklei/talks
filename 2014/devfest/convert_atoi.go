package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	// START OMIT
	/**
	 * int i = Integer.parseInt("42"); // throws NumberFormatException
	 *
	 **/
	i, err := strconv.Atoi("42")
	if err != nil {
		log.Fatal("not a integer")
	}

	// END OMIT
	fmt.Println(i)
}
