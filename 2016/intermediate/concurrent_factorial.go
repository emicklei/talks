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
	result := make(chan int, 10)
	for i := 1; i <= cap(result); i++ {
		go compute(i, result)
	}
	for i := 1; i <= cap(result); i++ {
		fmt.Println(<-result)
	}
	close(result)
}

// END OMIT
