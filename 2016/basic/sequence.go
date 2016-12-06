package main

import (
	"fmt"
	"sync"
)

//START OMIT
type Sequence struct {
	value int32
	mutex *sync.Mutex
}

// Next increments and returns the new value
func (s *Sequence) Next() int32 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.value += 1
	return s.value
}

//END OMIT
func main() {
	s := &Sequence{mutex: new(sync.Mutex)}
	fmt.Println(s.Next())
}
