package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	persons := []Person{}
	// .. which is the short version of
	// persons := make([]Person, 0)

	persons = append(persons, Person{Name: "Bob"})
	persons = append(persons, Person{Name: "Lisa"})

	for i, each := range persons {
		fmt.Println(i, each)
	}
}
