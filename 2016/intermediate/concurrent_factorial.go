package main

import (
	"fmt"
)

// START OMIT
func fac(i int) int {
	if i == 1 {
		return 1
	} else {
		return i * fac(i-1)
	}
}

func compute(i int, c chan int) {
	c <- fac(i)
}

func main() {
	result := make(chan int)
	for i := 1; i <= 10; i++ {
		go compute(i, result)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(<-result)
	}
}

// END OMIT
