package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var slides = `
intro.slide
embedding.slide
empty_interface.slide
interface.slide
pointer_receiver.slide
type_assertion.slide
`

func main() {
	writer, _ := os.Create("basic.slide")
	scanner := bufio.NewScanner(strings.NewReader(slides))
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			include(writer, line)
		}
	}
	writer.Close()
}

func include(w io.Writer, name string) {
	fmt.Println("including", name)
	file, _ := os.Open(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); line != "T" {
			io.WriteString(w, line)
			io.WriteString(w, "\n")
		}
	}
}
