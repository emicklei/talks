package main

// https://github.com/radovskyb/watcher/tree/master/cmd/watcher
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

var slides = `
intro.slide

architecture.slide

graphql.slide

query.slide

schema.slide

principles.slide

package.slide

selection.slide

codegen.slide
graphql-go.slide
codegen2.slide

gqlgen.slide
gqlc.slide
codegen3.slide

query.slide

resolver.slide

showmecode.slide
`

var slideCount = 0

func main() {
	flag.Parse()
	writer, _ := os.Create("main.slide")
	scanner := bufio.NewScanner(strings.NewReader(slides))
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			slideCount++
			include(writer, line)
		}
	}
	writer.Close()
}

func include(w io.Writer, name string) {
	fmt.Println("including", name)
	file, err := os.Open(strings.Trim(name, " "))
	if err != nil {
		log.Println(err.Error())
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
	if *debug {
		io.WriteString(w, fmt.Sprintf("_%s_ %d\n\n", name, slideCount))
	}
}
