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

func main() {
	now = func() time.Time {
		return time.Time{}
	}
	fmt.Println(NewMeeting("ik", now().Add(1*time.Hour)))
}

// END OMIT
