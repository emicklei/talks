package main

import "fmt"

func whatIsYourType(v interface{}) string {
	if s, ok := v.(string); ok {
		return fmt.Sprintf("%s is a string", s)
	}
	return "is something I don't know"
}

func main() {
	fmt.Println(whatIsYourType("hello"))
}
