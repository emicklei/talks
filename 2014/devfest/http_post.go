package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// START OMIT
	url := "http://www.bol.com/nl/catalog/browse_and_search/plain_list_page.html&Ntt=devfest"

	resp, _ := http.Post(url, "text/plain", nil)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// END OMIT
}
