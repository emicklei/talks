package main

import (
	"log"

	"github.com/cactus/go-statsd-client/statsd"
)

func main() {
	client, err := statsd.Dial("bnlimon6.local.nl.bol.com:18125", "prefix")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	client.Inc("ernest.go", 1, 1.0)

}
