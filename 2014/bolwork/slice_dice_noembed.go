package main

import "fmt"

//START OMIT
type stack []string

func (s *stack) push(entry string) {
	*s = append(*s, entry)
}

func (s *stack) pop() string {
	if len(*s) == 0 {
		panic("empty")
	}
	old := *s
	top := old[len(old)-1]
	*s = old[:len(old)-1]
	return top
}

func main() {
	s := stack{}
	s.push("hi")
	s.push("gopher")
	fmt.Println(s.pop(), s.pop())
}

//END OMIT
