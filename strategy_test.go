package retry

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFixedStrategy_NextDelay(t *testing.T) {
	strategy := &FixedStrategy{
		Interval:    100 * time.Millisecond,
		MaxAttempts: 3,
	}

	tests := []struct {
		attempt      int
		expectedWait time.Duration
		expectedOk   bool
	}{
		{0, 100 * time.Millisecond, true},
		{1, 100 * time.Millisecond, true},
		{2, 100 * time.Millisecond, true},
		{3, 0, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Attempt_%d", tt.attempt), func(t *testing.T) {
			delay, ok := strategy.NextDelay(tt.attempt)
			assert.Equal(t, tt.expectedWait, delay)
			assert.Equal(t, tt.expectedOk, ok)
		})
	}
}

func TestExponentialStrategy_NextDelay(t *testing.T) {
	strategy := &ExponentialStrategy{
		InitialInterval: 100 * time.Millisecond,
		MaxInterval:     800 * time.Millisecond,
		Multiplier:      2,
		MaxAttempts:     5,
	}

	tests := []struct {
		attempt      int
		expectedWait time.Duration
		expectedOk   bool
	}{
		{0, 100 * time.Millisecond, true},
		{1, 200 * time.Millisecond, true},
		{2, 400 * time.Millisecond, true},
		{3, 800 * time.Millisecond, true},
		{4, 800 * time.Millisecond, true}, // capped at MaxInterval
		{5, 0, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Attempt_%d", tt.attempt), func(t *testing.T) {
			delay, ok := strategy.NextDelay(tt.attempt)
			assert.Equal(t, tt.expectedWait, delay)
			assert.Equal(t, tt.expectedOk, ok)
		})
	}
}

func TestFibonacciStrategy_NextDelay(t *testing.T) {
	strategy := &FibonacciStrategy{
		InitialInterval: 100 * time.Millisecond,
		MaxInterval:     500 * time.Millisecond,
		MaxAttempts:     5,
	}

	tests := []struct {
		attempt      int
		expectedWait time.Duration
		expectedOk   bool
	}{
		{0, 100 * time.Millisecond, true},
		{1, 100 * time.Millisecond, true},
		{2, 200 * time.Millisecond, true},
		{3, 300 * time.Millisecond, true},
		{4, 500 * time.Millisecond, true}, // capped at MaxInterval
		{5, 0, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Attempt_%d", tt.attempt), func(t *testing.T) {
			delay, ok := strategy.NextDelay(tt.attempt)
			assert.Equal(t, tt.expectedWait, delay)
			assert.Equal(t, tt.expectedOk, ok)
		})
	}
}
