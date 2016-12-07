package main

import (
	"fmt"
	"strings"
)

func concat(names ...string) string {
	return join(" ", names...)
}

func join(separator string, words ...string) string {
	return strings.Join(words, separator)
}

//START OMIT
func main() {
	fmt.Printf("%q", concat("welcome", "to", "westworld"))
}

//END OMIT
