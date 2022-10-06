package main

// go install github.com/radovskyb/watcher/cmd/watcher@latest
// watcher -cmd="go run main.go"

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var debug = flag.Bool("d", false, "if true then write debug info")
var output = flag.String("o", "main.md", "merged Markdown file location")

// names starting with ! will be skipped
var slides = `
head.md
title.md
intro.md
whatisit.md

!overview.md 
note.md
!note_examples.md
sequence.md 
other_creates.md
composition.md 
drum.md
drum_merge.md 

nosound.md
midi_com.md 

demo.md
source_bits.md
`

var slideCount = 0

func main() {
	flag.Parse()
	writer, _ := os.Create(*output)
	scanner := bufio.NewScanner(strings.NewReader(slides))
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 && !strings.HasPrefix(line, "!") {
			slideCount++
			include(writer, line)
		}
	}
	writer.Close()
}

func include(w io.Writer, name string) {
	io.WriteString(w, "---\n")
	fmt.Println("including", name)
	file, err := os.Open(strings.Trim(name, " "))
	if err != nil {
		log.Println(err.Error())
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line := scanner.Text(); strings.Trim(line, " ") != "T" {
			io.WriteString(w, line)
			io.WriteString(w, "\n")
		}
	}
	io.WriteString(w, "\n")
	if *debug {
		io.WriteString(w, fmt.Sprintf("_%s_ %d\n\n", name, slideCount))
	}
}
