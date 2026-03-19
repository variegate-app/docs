package stopping_short

import (
	stdctx "context"
	"fmt"
	"sync/atomic"
	"time"
)

func Run(ctx stdctx.Context) error {
	// Generator stops when done is closed.
	done := make(chan struct{})
	var produced atomic.Int64

	out := make(chan int, 1)
	terminated := make(chan struct{})

	go func() {
		defer close(terminated)
		defer close(out)
		for i := 0; ; i++ {
			select {
			case <-done:
				return
			case <-ctx.Done():
				return
			case out <- i:
				produced.Add(1)
			}
		}
	}()

	// Consumer takes first N and then cancels the generator.
	const n = 5
	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("stopping-short: canceled: %w", ctx.Err())
		case <-done:
			return fmt.Errorf("stopping-short: done closed unexpectedly")
		case <-time.After(200 * time.Millisecond):
			return fmt.Errorf("stopping-short: timeout waiting for value")
		case <-out:
		}
	}

	close(done)

	select {
	case <-terminated:
	case <-ctx.Done():
		return fmt.Errorf("stopping-short: generator did not terminate: %w", ctx.Err())
	}

	after := produced.Load()
	time.Sleep(20 * time.Millisecond)
	if produced.Load() != after {
		return fmt.Errorf("stopping-short: expected no more production after done close, before=%d after=%d", after, produced.Load())
	}
	return nil
}

