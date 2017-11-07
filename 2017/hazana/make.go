package main

// https://github.com/radovskyb/watcher/tree/master/cmd/watcher
// watcher -cmd="go run make.go"

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"log" 
	"strings"
)

var slides = `
intro.slide
grpc.slide
load_profile.slide
rampup.slide

proto.slide
server.slide
client.slide

profile_ideal.slide
profile_rampup.slide
profile_goroutines.slide
rampup_strategy.slide

attack_interface.slide
attack_interface2.slide

loadrun.slide
loadrun_main.slide
`

/**
server
client
loadprofile
attacker
hazana reporting
run in the clouds
hazana docker
**/

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
}
