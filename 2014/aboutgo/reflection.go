package main

import "encoding/xml"

type Product struct {
	Id    string
	Title string
}

var doc = `<?xml version="1.0" ?>
<Product>
	<Id>100043432</Id>
	<Title>Life of Pi</Title>
</Product>`

func main() {
	p := new(Product)
	xml.Unmarshal([]byte(doc), p) // HL
	println(p.Title)
}
