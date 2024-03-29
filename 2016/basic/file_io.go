package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	echoLines("intro.slide")
}

// START OMIT
func echoLines(name string) {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// END OMIT
