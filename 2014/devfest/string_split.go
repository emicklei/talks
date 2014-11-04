package main

import (
	"fmt"
	"strings"
)

func main() {

	// START OMIT
	/**
	 *
	 * String host = "10.10.2.63:8080".split(":")[0];
	 *
	 **/

	host := strings.Split("10.10.2.63:8080", ":")[0]
	// END OMIT
	fmt.Println(host)
}
