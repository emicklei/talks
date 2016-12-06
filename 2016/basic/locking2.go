package main

import (
	"fmt"
	"sync"
	"time"
)

type Map interface {
	Put(k string, v time.Time)
	Get(k string) (bool, time.Time)
}

type ConcurrentHashMap struct {
	mutex    *sync.RWMutex
	elements map[string]time.Time
}

func NewConcurrentHashMap() *ConcurrentHashMap {
	return &ConcurrentHashMap{
		mutex:    new(sync.RWMutex),
		elements: map[string]time.Time{},
	}
}

// START OMIT

func (c *ConcurrentHashMap) Put(k string, v time.Time) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.elements[k] = v
}

func (c *ConcurrentHashMap) Get(k string) (ok bool, value time.Time) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok = c.elements[k]
	return
}

func main() {
	var m Map = NewConcurrentHashMap()
	m.Put("now", time.Now())
	fmt.Println(m.Get("now"))
}

// END OMIT
