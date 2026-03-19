package fault_tolerance

import (
	stdctx "context"
	"fmt"
	"sync/atomic"
	"time"
)

func Run(ctx stdctx.Context) error {
	// Supervisor runs a worker, recovers panics, restarts it a few times.
	var attempts int64
	var recovered int64

	worker := func() {
		n := atomic.AddInt64(&attempts, 1)
		if n == 1 {
			panic("boom")
		}
		// second run succeeds
	}

	done := make(chan struct{})
	go func() {
		defer close(done)
		restarts := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						atomic.AddInt64(&recovered, 1)
					}
				}()
				worker()
			}()

			// stop after success (worker doesn't panic)
			if atomic.LoadInt64(&attempts) >= 2 {
				return
			}

			restarts++
			if restarts > 3 {
				return
			}
			time.Sleep(5 * time.Millisecond)
		}
	}()

	select {
	case <-done:
	case <-ctx.Done():
		return fmt.Errorf("fault-tolerance: canceled: %w", ctx.Err())
	}

	if atomic.LoadInt64(&attempts) < 2 {
		return fmt.Errorf("fault-tolerance: expected at least 2 attempts, got %d", attempts)
	}
	if atomic.LoadInt64(&recovered) < 1 {
		return fmt.Errorf("fault-tolerance: expected panic to be recovered at least once, got %d", recovered)
	}
	return nil
}

