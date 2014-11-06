package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// START OMIT

	file, err := os.Open("2014/devfest/quotes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
