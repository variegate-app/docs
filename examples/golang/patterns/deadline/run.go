package deadline

import (
	stdctx "context"
	"errors"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	deadlineCtx, cancel := stdctx.WithDeadline(ctx, time.Now().Add(30*time.Millisecond))
	defer cancel()

	select {
	case <-time.After(50 * time.Millisecond):
		// Not expected to happen fast enough.
		return errors.New("deadline: unexpected completion")
	case <-deadlineCtx.Done():
		fmt.Println("deadline: timed out as expected")
		return nil
	}
}

