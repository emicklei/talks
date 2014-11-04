package main

import (
	"fmt"
	"sort"
)

func main() {

	// START OMIT
	/**
	 * String[] names = new String[]{"olivia","melissa","judy","rose"};
	 * Arrays.sort(names);
	 *
	 **/

	names := []string{"erik", "gerard", "hugo", "karel", "neo", "bruce"}

	sort.StringSlice(names).Sort()

	// END OMIT
	fmt.Println(names)
}
