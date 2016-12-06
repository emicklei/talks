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
func echoLines(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		fmt.Println(scanner.Text())
	}
	return nil
}

// END OMIT
