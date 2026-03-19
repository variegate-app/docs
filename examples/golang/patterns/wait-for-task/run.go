package wait_for_task

import (
	stdctx "context"
	"fmt"
)

func Run(ctx stdctx.Context) error {
	ch := make(chan string)
	terminated := make(chan struct{})
	errc := make(chan error, 1)

	go func() {
		defer close(terminated)
		v := <-ch
		if v != "paper" {
			errc <- fmt.Errorf("wait-for-task: expected paper, got %q", v)
			return
		}
		errc <- nil
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("wait-for-task: canceled: %w", ctx.Err())
	case ch <- "paper":
	}

	select {
	case <-terminated:
		return <-errc
	case <-ctx.Done():
		return fmt.Errorf("wait-for-task: worker did not terminate: %w", ctx.Err())
	}
}

