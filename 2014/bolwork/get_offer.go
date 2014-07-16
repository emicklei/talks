package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//START OMIT
func main() {
	pos := http.Client{}

	req, err := http.NewRequest("GET",
		"http://productoffer.services.test2.bol.com/slo/products/9000000011602483/offer", nil)
	if err != nil {
		panic("request build failed")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := pos.Do(req)
	if err != nil {
		panic("request send failed")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("response read failed")
	}

	fmt.Println(string(body))
}

//END OMIT
