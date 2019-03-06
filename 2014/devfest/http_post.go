package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// START OMIT
	url := "https://www.bol.com/nl/p/harry-potter-boxset/9200000023479731"

	resp, _ := http.Post(url, "text/plain", nil)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// END OMIT
}
