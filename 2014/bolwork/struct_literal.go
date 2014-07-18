package main

import "fmt"

type Rectangle struct {
	Left, Right, Width, Height int
}

func main() {
	box := Rectangle{
		Left:   0,
		Right:  0,
		Width:  1024,
		Height: 768,
	}
	fmt.Println(box)
}
