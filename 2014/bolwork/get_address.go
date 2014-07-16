package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println(http.Get("http://localhost:7777/books"))
}
