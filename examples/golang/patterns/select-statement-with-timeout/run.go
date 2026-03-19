package select_statement_with_timeout

import (
	stdctx "context"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	// Channel never gets a value => select must hit timeout case.
	ch := make(chan string)

	select {
	case <-ctx.Done():
		return fmt.Errorf("select-statement-with-timeout: canceled early: %w", ctx.Err())
	case <-time.After(30 * time.Millisecond):
		// expected
		return nil
	case v := <-ch:
		_ = v
		return fmt.Errorf("select-statement-with-timeout: unexpected value received")
	}
}

