package retry

import "time"

func Do(fn func() error, strategy Strategy) error {
	attempt := 0
	for {
		err := fn()
		if err == nil {
			return nil
		}

		delay, shouldRetry := strategy.NextDelay(attempt)
		if !shouldRetry {
			return err
		}

		time.Sleep(delay)
		attempt++
	}
}

func MustDo(fn func() error, strategy Strategy) {
	if err := Do(fn, strategy); err != nil {
		panic(err)
	}
}
