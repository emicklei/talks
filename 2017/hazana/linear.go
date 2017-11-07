func (s linearIncreasingGoroutinesAndRequestsPerSecondStrategy) execute(r *runner) {
	r.spawnAttacker() // start at least one
	for i := 1; i <= r.config.RampupTimeSec; i++ {
		spawnAttackersToSize(r, i*r.config.MaxAttackers/r.config.RampupTimeSec)
		takeDuringOneRampupSecond(r, i)
	}
}