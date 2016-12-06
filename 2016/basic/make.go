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

varargs.slide 

embedding.slide

pointer_receiver.slide

empty_interface.slide
interface.slide

link_slice_dice.slide

locking.slide
http_server.slide 

func_value.slide
time_inject.slide

file_io.slide

api_interface.slide

got_want.slide
table_driven.slide

closing_notes.slide
`

// error.slide
// type_assertion.slide

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
