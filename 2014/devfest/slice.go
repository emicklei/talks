package main

import "fmt"

func main() {

	// START OMIT
	/**
	 *
	 * List<String> colors1 = Arrays.asList("red", "green", "blue");
	 * colors1.add("black");
	 *
	 * String[] colors2 = new String[]{"red", "green", "blue"};
	 *
	 **/

	colors1 := []string{"red", "green", "blue"}
	colors1 = append(colors1, "black")

	// var colors2 = [3]string
	//
	colors2 := [...]string{"red", "green", "blue"}
	// END OMIT
	fmt.Println(colors1)
	fmt.Println(colors2)
}
