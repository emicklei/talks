package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pos := http.Client{}

	req, _ := http.NewRequest("GET", "http://productoffer.services.test2.bol.com/slo/products/9000000011602483/offer", nil)

	req.Header.Set("Accept", "application/json")

	resp, _ := pos.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
