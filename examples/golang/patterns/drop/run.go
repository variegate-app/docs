package drop

import (
	stdctx "context"
	"fmt"
	"sync/atomic"
	"time"
)

func Run(ctx stdctx.Context) error {
	done := ctx.Done()

	// Capacity is the maximum number of in-flight items.
	const cap = 3
	ch := make(chan int, cap)
	start := make(chan struct{})

	var processed int64

	workerDone := make(chan struct{})
	go func() {
		defer close(workerDone)
		<-start
		for v := range ch {
			_ = v
			atomic.AddInt64(&processed, 1)
			// Slow consumer to keep buffer full during producer phase.
			select {
			case <-done:
				return
			case <-time.After(10 * time.Millisecond):
			}
		}
	}()

	input := 10
	for i := 0; i < input; i++ {
		select {
		case <-done:
			return fmt.Errorf("drop: canceled: %w", ctx.Err())
		default:
		}

		// Drop policy: if buffer full, discard item.
		select {
		case ch <- i:
		default:
			// dropped
		}
	}

	close(start) // now the worker starts consuming
	close(ch)    // stop after producer phase

	select {
	case <-workerDone:
	case <-done:
		return fmt.Errorf("drop: timeout: %w", ctx.Err())
	}

	p := int(atomic.LoadInt64(&processed))
	// With deterministic ordering and cap=3, we expect to process <= cap.
	if p <= 0 || p > cap {
		return fmt.Errorf("drop: expected processed in 1..%d, got %d", cap, p)
	}
	// Must have dropped something since input > cap and consumer starts after producer.
	if p == input {
		return fmt.Errorf("drop: expected some items dropped, but processed all %d", input)
	}
	return nil
}

