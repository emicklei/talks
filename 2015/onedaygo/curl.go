package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		log.Prinln("Usage: curl [url]")
		return
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
