package wait_for_result

import (
	stdctx "context"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	// Signal channel: worker sends value, manager waits.
	ch := make(chan string, 1)

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-time.After(10 * time.Millisecond):
		}
		ch <- "paper"
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("wait-for-result: canceled: %w", ctx.Err())
	case v := <-ch:
		if v != "paper" {
			return fmt.Errorf("wait-for-result: expected paper, got %q", v)
		}
		return nil
	}
}

