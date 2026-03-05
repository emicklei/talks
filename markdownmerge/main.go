package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var debug = flag.Bool("d", false, "if true then write debug info")

// names starting with ! will be skipped
var input = flag.String("i", "slides.txt", "file with list of Markdown files to merge")
var output = flag.String("o", "main.md", "merged Markdown file location")

var slideCount = 0

func main() {
	flag.Parse()
	reader, err := os.Open(*input)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	writer, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	dir := filepath.Dir(*input)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 && !strings.HasPrefix(line, "!") {
			slideCount++
			include(writer, filepath.Join(dir, line))
		} else {
			fmt.Println("skipping", line)
		}
	}
}

func include(w io.Writer, name string) {
	if slideCount > 1 {
		io.WriteString(w, "---\n")
	}
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
