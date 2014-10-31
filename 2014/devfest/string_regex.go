package main

import (
	"fmt"
	"regexp"
)

func main() {

	// START OMIT
	/**
	 * String top = "10.10.2.63:8080".split("[\\.:]")[0];
	**/
	re := regexp.MustCompile("[\\.:]")
	top := re.Split("10.10.2.63:8080", -1)[0]
	// END OMIT
	fmt.Println(top)
}
