package main

import "fmt"

// START OMIT
func useTypeAssertion(v interface{}) string {
	if s, ok := v.(string); ok {
		return fmt.Sprintf("%s is a string", s)
	}
	return "is something I don't know"
}

func useTypeSwitch(v interface{}) string {
	switch v.(type) {
	case string:
		return fmt.Sprintf("%v is a string", v)
	case float64:
		return fmt.Sprintf("%v is a float 64", v)
	default:
		return "is something I don't know"
	}
}

// END OMIT
func main() {
	fmt.Println(useTypeAssertion("hello"))
	fmt.Println(useTypeSwitch(3.14159))
}
