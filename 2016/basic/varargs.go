package main

import "fmt"

func concat(names ...string) string {
	return join(" ", names...)
}

func join(separator string, words ...string) string {
	return "TODO"
}

//START OMIT
func main() {
	fmt.Println(concat("welcome", "to", "westworld"))
}

//END OMIT
