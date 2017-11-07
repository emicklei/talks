func (s spawnAsWeNeedStrategy) execute(r *runner) {
	r.spawnAttacker() // start at least one
	for i := 1; i <= r.config.RampupTimeSec; i++ {
		targetRate, lastMetrics := takeDuringOneRampupSecond(r, i)
		currentRate := lastMetrics.Rate
		if currentRate < float64(targetRate) {
			factor := float64(targetRate) / currentRate
			if factor > 2.0 {
				factor = 2.0
			}
			spawnAttackersToSize(r, int(math.Ceil(float64(len(r.attackers))*factor)))
		}
	}
}
