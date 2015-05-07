package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// START OMIT
	url := "http://www.bol.com/nl/catalog/browse_and_search/plain_list_page.html&Ntt=devfest"
	var buf bytes.Buffer
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept-Language", "nl,en;q=0.5")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	// END OMIT
}
