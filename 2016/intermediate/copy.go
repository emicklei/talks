package main

import (
	"io"
	"log"
	"os"

	"github.com/emicklei/tre"
)

func main() {
	if len(os.Args) != 3 {
		println("Usage copy <from> <to>")
		return
	}
	if err := copyFile(os.Args[1], os.Args[2]); err != nil {
		log.Println("problem copying file", err)
	}
}

// START OMIT
func copyFile(from, to string) error {
	input, err := os.Open(from)
	if err != nil {
		return tre.New(err, "cannot open file for read", "from", from, "err", err)
	}
	// make explict that I do not care about this error
	defer func() { _ = input.Close() }()

	output, err := os.Create(to)
	if err != nil {
		return tre.New(err, "cannot open file for write", "to", to, "err", err)
	}
	// what developers usually write
	defer output.Close()

	if _, err := io.Copy(output, input); err != nil {
		return tre.New(err, "cannot copy file", "from", from, "to", to, "err", err)
	}

	return nil
}

// END OMIT
