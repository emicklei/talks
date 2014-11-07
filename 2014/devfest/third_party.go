package main

import (
	"fmt"
	"log"

	"github.com/SlyMarbo/rss"
)

func main() {
	feed, err := rss.Fetch("http://www.nu.nl/feeds/rss/algemeen.rss")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", feed)
}
