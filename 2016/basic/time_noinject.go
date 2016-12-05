package main

import "time"

// START OMIT
type Meeting struct {
	Created time.Time
	When    time.Time
	Who     string
}

func NewMeeting(who string, when time.Time) Meeting {
	return Meeting{Created: time.Now(), When: when, Who: who}
}

// END OMIT
