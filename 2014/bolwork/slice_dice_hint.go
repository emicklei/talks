package main

import "fmt"

func main() {
	s := newStack()
	s.push("hi")
	s.push("gopher")
	fmt.Println(s.pop(), s.pop())
}
