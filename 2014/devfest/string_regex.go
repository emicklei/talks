package main

import (
	"fmt"
	"regexp"
)

func main() {

	// START OMIT
	/**
	 *
	 * String port = "10.10.2.63:8080".split("[\\.:]")[4];
	 *
	 **/

	re := regexp.MustCompile("[\\.:]")
	port := re.Split("10.10.2.63:8080", -1)[4]
	// END OMIT
	fmt.Println(port)
}
