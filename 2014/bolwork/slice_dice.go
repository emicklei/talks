package main

import "fmt"

//START OMIT
type stack struct {
	elements []string
}

func newStack() stack {
	return stack{}
}

func (s *stack) push(entry string) {
	s.elements = append(s.elements, entry)
}

func (s *stack) pop() string {
	if len(s.elements) == 0 {
		panic("empty")
	}
	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return top
}

//END OMIT

func main() {
	s := newStack()
	s.push("hi")
	s.push("gopher")
	fmt.Println(s.pop(), s.pop())
}
