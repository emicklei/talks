package main

import "fmt"

type stack struct {
	// ...
}

func main() {
	s := newStack()
	s.push("hi")
	s.push("gopher")
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
