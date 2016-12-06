package main

import (
	"bufio"
	"fmt"
	"os"
)

// START OMIT
func main() {
	if lines, err := readLines("basic.slide"); err == nil {
		writeLines(lines, "/tmp/cisab.slide")
	}
}

func readLines(name string) (lines []string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return
		}
		lines = append(lines, scanner.Text())
	}
	return
}

// END OMIT
func writeLines(lines []string, name string) (err error) {
	file, err := os.Create(name)
	if err != nil {
		return
	}
	defer file.Close()
	for i := len(lines) - 1; i > 0; i-- {
		fmt.Fprintf(file, "%s\n", reverse(lines[i]))
	}
	return
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
