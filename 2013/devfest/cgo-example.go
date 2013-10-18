package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(2)
	fmt.Printf("C random:%d \n", Random())
}
