package for_select

import (
	stdctx "context"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c1 <- 1
	c2 <- 2
	close(c1)
	close(c2)

	got := make(map[int]bool, 2)
	for len(got) < 2 {
		select {
		case <-ctx.Done():
			return fmt.Errorf("for-select: canceled: %w", ctx.Err())
		case v, ok := <-c1:
			if !ok {
				c1 = nil
				continue
			}
			got[v] = true
		case v, ok := <-c2:
			if !ok {
				c2 = nil
				continue
			}
			got[v] = true
		default:
			// avoid busy loop in this demo
			select {
			case <-ctx.Done():
				return fmt.Errorf("for-select: canceled: %w", ctx.Err())
			case <-time.After(1 * time.Millisecond):
			}
		}
	}

	if !got[1] || !got[2] {
		return fmt.Errorf("for-select: expected values {1,2}, got=%v", got)
	}
	return nil
}

