package main

import (
	"fmt"
	"sync"
)

//START OMIT
type Sequence struct {
	sync.Mutex
	value int32
}

// Next increments and returns the new value
func (s *Sequence) Next() int32 {
	s.Lock()
	defer s.Unlock()
	s.value += 1
	return s.value
}

//END OMIT
func main() {
	s := new(Sequence)
	fmt.Println(s.Next())
}
