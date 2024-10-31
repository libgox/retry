package retry

import "time"

type Strategy interface {
	NextDelay(attempt int) (time.Duration, bool)
}

type FixedStrategy struct {
	Interval    time.Duration
	MaxAttempts int
}

func (f *FixedStrategy) NextDelay(attempt int) (time.Duration, bool) {
	if f.MaxAttempts != -1 && attempt >= f.MaxAttempts {
		return 0, false
	}
	return f.Interval, true
}

type ExponentialStrategy struct {
	InitialInterval time.Duration
	MaxInterval     time.Duration
	Multiplier      float64
	MaxAttempts     int

	currentDelay time.Duration
}

func (e *ExponentialStrategy) NextDelay(attempt int) (time.Duration, bool) {
	if e.MaxAttempts != -1 && attempt >= e.MaxAttempts {
		return 0, false
	}

	if attempt == 0 {
		e.currentDelay = e.InitialInterval
		return e.currentDelay, true
	}

	nextDelay := time.Duration(float64(e.currentDelay) * e.Multiplier)
	if nextDelay > e.MaxInterval || nextDelay <= 0 {
		nextDelay = e.MaxInterval
	}

	e.currentDelay = nextDelay
	return e.currentDelay, true
}

type FibonacciStrategy struct {
	InitialInterval time.Duration
	MaxInterval     time.Duration
	MaxAttempts     int

	previousDelay time.Duration
	currentDelay  time.Duration
}

func (f *FibonacciStrategy) NextDelay(attempt int) (time.Duration, bool) {
	if f.MaxAttempts != -1 && attempt >= f.MaxAttempts {
		return 0, false
	}

	if attempt == 0 {
		f.previousDelay = 0
		f.currentDelay = f.InitialInterval
		return f.currentDelay, true
	}

	nextDelay := f.previousDelay + f.currentDelay
	if nextDelay > f.MaxInterval || nextDelay <= 0 {
		nextDelay = f.MaxInterval
	}

	f.previousDelay = f.currentDelay
	f.currentDelay = nextDelay
	return f.currentDelay, true
}
