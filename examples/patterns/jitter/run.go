package jitter

import (
	stdctx "context"
	"fmt"
	"math/rand"
	"time"
)

func FullJitter(base time.Duration) time.Duration {
	if base <= 0 {
		return 0
	}
	return time.Duration(rand.Int63n(int64(base)))
}

func EqualJitter(min, max time.Duration) time.Duration {
	if max <= min {
		return min
	}
	half := (max - min) / 2
	if half <= 0 {
		return min
	}
	return min + time.Duration(rand.Int63n(int64(half)))
}

func Run(ctx stdctx.Context) error {
	rand.Seed(time.Now().UnixNano())

	base := 100 * time.Millisecond
	min := 50 * time.Millisecond
	max := 150 * time.Millisecond

	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("jitter: canceled: %w", ctx.Err())
		default:
		}

		v := FullJitter(base)
		if v < 0 || v >= base {
			return fmt.Errorf("jitter: FullJitter out of range: got=%v base=%v", v, base)
		}

		e := EqualJitter(min, max)
		// In our implementation, EqualJitter returns in [min, min+half).
		if e < min || e >= max {
			return fmt.Errorf("jitter: EqualJitter out of range: got=%v min=%v max=%v", e, min, max)
		}
	}
	return nil
}

