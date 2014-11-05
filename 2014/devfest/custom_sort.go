package main

import (
	"fmt"
	"sort"
)

type City struct {
	Name  string
	Count int
}

var cities = []City{
	City{"Utrecht", 330772},
	City{"Leons", 22},
	City{"Amersfoort", 151534},
	City{"Amsterdam", 813562},
}

// START OMIT
type CitySorter []City

func (s CitySorter) Len() int { // HL
	return len(s)
}
func (s CitySorter) Swap(i, j int) { // HL
	s[i], s[j] = s[j], s[i]
}
func (s CitySorter) Less(i, j int) bool { // HL
	return s[i].Count < s[j].Count
}

func main() {
	sort.Sort(sort.Reverse(CitySorter(cities)))

	fmt.Println(cities)
}

// END OMIT
