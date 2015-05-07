package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// START OMIT
func main() {
	if len(os.Args) == 0 {
		log.Println("Usage: curl [url]")
		return
	}
	cmd := exec.Command("curl", os.Args[1])
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("unable to execute curl because:%v", err)
	}
	fmt.Fprintln(os.Stdout, string(data))
}

// END OMIT
