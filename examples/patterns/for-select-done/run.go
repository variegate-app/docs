package for_select_done

import (
	stdctx "context"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	done := make(chan struct{})
	terminated := make(chan struct{})

	go func() {
		defer close(terminated)
		for {
			select {
			case <-done:
				return
			default:
				// simulate work without blocking forever
				select {
				case <-ctx.Done():
					return
				case <-time.After(1 * time.Millisecond):
				}
			}
		}
	}()

	close(done)

	select {
	case <-terminated:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("for-select-done: worker did not stop: %w", ctx.Err())
	}
}

