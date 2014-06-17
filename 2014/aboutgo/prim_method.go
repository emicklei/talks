package main

import "fmt"

// START OMIT
type YesNo bool

func (y YesNo) String() string {
	if y {
		return "yes"
	} else {
		return "no"
	}

}

func main() {
	fmt.Println(YesNo(true))
}

// END OMIT
