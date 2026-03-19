package steady_state

import (
	stdctx "context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Run(ctx stdctx.Context) error {
	// Steady-state = bounded queue + limited worker pool.
	const (
		workers = 2
		queueCap = 5
		n        = 20
	)

	jobs := make(chan int, queueCap)
	var processed atomic.Int64

	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for j := range jobs {
				_ = j
				processed.Add(1)
				time.Sleep(2 * time.Millisecond)
			}
		}()
	}

	// Track maximum observed queue length (safe in main goroutine).
	maxLen := 0

	for i := 0; i < n; i++ {
		select {
		case <-ctx.Done():
			close(jobs)
			wg.Wait()
			return fmt.Errorf("steady-state: canceled: %w", ctx.Err())
		case jobs <- i:
			if l := len(jobs); l > maxLen {
				maxLen = l
			}
		}
	}
	close(jobs)
	wg.Wait()

	if processed.Load() != n {
		return fmt.Errorf("steady-state: expected processed=%d got=%d", n, processed.Load())
	}
	if maxLen > queueCap {
		return fmt.Errorf("steady-state: expected maxLen<=%d got %d", queueCap, maxLen)
	}
	return nil
}

