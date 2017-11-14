func attack(attacker Attack, next, quit <-chan bool, results chan<- result, timeout time.Duration) {
	for {
		select {
		case <-next:
			begin := time.Now()
			done := make(chan DoResult)
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			go func() {
				done <- attacker.Do(ctx)
			}()
			var dor DoResult
			// either get the result from the attacker or from the timeout
			select {
			case <-ctx.Done():
				dor = DoResult{Error: errAttackDoTimedOut}
			case dor = <-done:
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