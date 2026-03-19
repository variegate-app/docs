package explicit_cancellation

import (
	stdctx "context"
	"fmt"
	"sync"
	"time"
)

func Run(ctx stdctx.Context) error {
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		t := time.NewTicker(10 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-done:
				return
			case <-t.C:
				// keep working
			}
		}
	}()

	// Cancel explicitly and wait for the worker to exit.
	close(done)
	wg.Wait()
	fmt.Println("explicit-cancellation: worker stopped on closed done channel")
	return nil
}

