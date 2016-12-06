package main

import (
	"fmt"
	"strconv"
	"sync"
)

func fetchOffer(id int) string {
	return strconv.Itoa(id)
}

//START OMIT
func main() {
	ids := []int{9000000065433456, 800000005435343, 7000000011602483, 600000005432543}

	wg := new(sync.WaitGroup)

	for _, each := range ids {
		wg.Add(1)

		go func() {
			fmt.Println(fetchOffer(each)) // do you see the BUG?
			wg.Done()
		}()
	}

	wg.Wait() // for all started goroutines to finish
}

//END OMIT
