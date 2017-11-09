func attack(attacker Attack, next, quit <-chan bool, results chan<- result, timeout time.Duration) {
	for {
		select {
		case <-next:
			begin := time.Now()
			done := make(chan DoResult)
			go func() {
				done <- attacker.Do() // HL
			}()
			var dor DoResult
			select {
			case <-time.After(timeout): // HL
				dor = DoResult{Error: errAttackDoTimedOut}
			case dor = <-done: // HL
			}
			end := time.Now()
			results <- result{
				doResult: dor,
				begin:    begin,
				end:      end,
				elapsed:  end.Sub(begin),
			}
		case <-quit:
			return
		}
	}
}