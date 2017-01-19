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

goroutine.slide

json_http.slide

error_handling.slide

closing_notes.slide
`

// https://github.com/cespare/reflex
// reflex -r '\.slide$' go run make.go

func main() {
	writer, _ := os.Create("main.slide")
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
	file, err := os.Open(strings.Trim(name, " "))
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); strings.Trim(line, " ") != "T" {
			io.WriteString(w, line)
			io.WriteString(w, "\n")
		}
	}
}
