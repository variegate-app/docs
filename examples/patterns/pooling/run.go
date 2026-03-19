package pooling

import (
	stdctx "context"
	"fmt"
	"sync"
	"sync/atomic"
)

func Run(ctx stdctx.Context) error {
	const workers = 3
	const jobs = 10

	// Unbuffered signal channel: when manager sends a signal, exactly one worker receives it.
	ch := make(chan int)
	var processed atomic.Int64

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					_ = v
					processed.Add(1)
				}
			}
		}()
	}

	for i := 0; i < jobs; i++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("pooling: canceled: %w", ctx.Err())
		case ch <- i:
		}
	}

	close(ch) // stop workers
	wg.Wait()

	if processed.Load() != jobs {
		return fmt.Errorf("pooling: expected %d processed, got %d", jobs, processed.Load())
	}
	return nil
}

