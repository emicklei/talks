func takeDuringOneRampupSecond(r *runner, second int) (int, *Metrics) {
	// ... code omitted

	// for each second start a new reduced rate limiter
	rps := second * r.config.RPS / r.config.RampupTimeSec
	if rps == 0 { // minimal 1
		rps = 1
	}
	limiter := ratelimit.New(rps)
	oneSecondAhead := time.Now().Add(time.Duration(1 * time.Second))
	// put the attackers to work
	for time.Now().Before(oneSecondAhead) {
		limiter.Take()
		r.next <- true
	}
		
	// ... code omitted
	return rps, rampMetrics
}
