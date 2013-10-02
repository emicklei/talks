package main

import "testing"

func TestSanity(t *testing.T) {
	if 1+1 != 2 {
		t.Error("uhh...")
	}
}

func main() {
	tests := []testing.InternalTest{{"TestSanity", TestSanity}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}
