package context

import (
	stdctx "context"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	ctx, cancel := stdctx.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	// Simulate “external call” that completes before timeout.
	select {
	case <-time.After(10 * time.Millisecond):
		fmt.Println("context: ok (completed before timeout)")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

