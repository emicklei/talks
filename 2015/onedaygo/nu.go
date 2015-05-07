package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// START OMIT
type RSS struct {
	Channel struct {
		Item []struct {
			Title string `xml:"title"`
		} `xml:"item"`
	} `xml:"channel"`
}

func main() {
	resp, err := http.Get("http://www.nu.nl/rss/Algemeen")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	rss := new(RSS)
	if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
		log.Fatal(err)
	}
	for _, each := range rss.Channel.Item {
		fmt.Println(each.Title)
	}
}

// END OMIT
