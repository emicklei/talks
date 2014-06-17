package main

// START OMIT
import "testing"

func TestSeven(t *testing.T) {
	seven := 3 + 4
	if seven != 7 {
		t.Errorf("huh? then what:%d", seven)
	}
}

// END OMIT

func main() {
	tests := []testing.InternalTest{{"TestSeven", TestSeven}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}
