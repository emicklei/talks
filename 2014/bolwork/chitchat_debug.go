package main

import (
	"fmt"

	"github.com/emicklei/hopwatch"
)

const (
	N = 20
)

//START OMIT
func factorial(n int) int {
	hopwatch.Dump(n).Break()
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	output := make(chan int)
	for i := 1; i < N; i++ {
		go func(n int) {
			output <- factorial(n)
		}(i)
	}
	for i := 1; i < N; i++ {
		fmt.Println(<-output)
	}
}

//END OMIT
