package main

import (
	"fmt"
	"time"
)

// START OMIT
var now = time.Now

type Meeting struct {
	Created time.Time
	When    time.Time
	Who     string
}

func NewMeeting(who string, when time.Time) Meeting {
	return Meeting{Created: now(), When: when, Who: who}
}

// this would be a test function instead
func main() {
	now = func() time.Time {
		// Mon Jan 2 15:04:05 -0700 MST 2006
		t, _ := time.Parse("2006-Jan-02", "2016-Dec-05")
		return t
	}
	fmt.Println(NewMeeting("boss", now().Add(24*time.Hour)))
}

// END OMIT
