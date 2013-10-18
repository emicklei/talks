package main

import (
	"io"
	"os"
)

// START OMIT

type CountingWriter struct {
	io.Writer
	count int
}

func (i CountingWriter) Write(p []byte) (n int, err error) {
	n, err = i.Writer.Write(p)
	i.count += n
	return
}

// END OMIT

func main() {
	cw := CountingWriter{os.Stdout}
	cw.Write([]byte{byte(0)})
}
