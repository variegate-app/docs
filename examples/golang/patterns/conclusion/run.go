package conclusion

import (
	stdctx "context"
	"fmt"
)

func Run(ctx stdctx.Context) error {
	// This package is a “meta” conclusion; correctness means the process finishes.
	select {
	case <-ctx.Done():
		return fmt.Errorf("conclusion: canceled: %w", ctx.Err())
	default:
		return nil
	}
}

