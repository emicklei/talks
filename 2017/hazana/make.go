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
	"flag"
)

var debug = flag.Bool("d",false,"if true then write debug info")

var slides = `
intro.slide
agenda.slide

grpc.slide


proto.slide
server.slide
client.slide

profile_ideal.slide
load_profile.slide
rampup.slide

hazana.slide
attack_interface.slide

loadrun.slide
loadrun_main.slide

demo.slide

runner.slide
config.slide
token_flow.slide

attack_inside.slide
profile_rampup.slide
rampup_inside.slide

profile_goroutines.slide
rampup_strategy.slide
linear.slide
exp2.slide

deploy.slide
docker.slide
docker_run.slide

wrapup.slide
`

var slideCount = 0

func main() {
	flag.Parse()
	writer, _ := os.Create("basic.slide")
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
		io.WriteString(w, fmt.Sprintf("_%s_ %d\n\n",name,slideCount))
	}
}
