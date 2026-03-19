package cancellation

import (
	stdctx "context"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	ctx, cancel := stdctx.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		// Simulate work that periodically checks cancellation.
		t := time.NewTicker(10 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				// keep working
			}
		}
	}()

	<-done
	fmt.Println("cancellation: worker stopped on ctx.Done()")
	return nil
}

