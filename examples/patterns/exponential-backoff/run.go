package exponential_backoff

import (
	stdctx "context"
	"fmt"
	"time"
)

func RetryWithBackoff(ctx stdctx.Context, maxRetries int, baseDelay time.Duration, fn func() error) error {
	var err error
	start := time.Now()
	for attempt := 0; attempt <= maxRetries; attempt++ {
		if err = fn(); err == nil {
			return nil
		}
		if attempt == maxRetries {
			break
		}

		delay := baseDelay * time.Duration(1<<attempt)
		// Backoff delay must respect cancellation.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
	}
	_ = start
	return err
}

func Run(ctx stdctx.Context) error {
	ctx, cancel := stdctx.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	const maxRetries = 5
	const base = 5 * time.Millisecond

	attempt := 0
	start := time.Now()
	err := RetryWithBackoff(ctx, maxRetries, base, func() error {
		attempt++
		if attempt < 3 {
			return fmt.Errorf("fail attempt=%d", attempt)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("exponential-backoff: expected success, got err=%v", err)
	}
	if attempt != 3 {
		return fmt.Errorf("exponential-backoff: expected 3 attempts, got %d", attempt)
	}
	// We expect at least baseDelay + baseDelay*2 of waiting (attempts-1 delays).
	minElapsed := base + 2*base
	if time.Since(start) < minElapsed-2*time.Millisecond {
		return fmt.Errorf("exponential-backoff: elapsed too small: got=%v expected at least %v", time.Since(start), minElapsed)
	}
	return nil
}

