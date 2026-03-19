package fail_fast

import (
	stdctx "context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Run(ctx stdctx.Context) error {
	// Fail-fast: cancel the whole group on the first error.
	ctx, cancel := stdctx.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	inputs := []string{"ok-1", "bad", "ok-2", "ok-3"}

	var canceled atomic.Int64

	doOne := func(in string) error {
		if in == "bad" {
			return errors.New("boom")
		}
		// Wait until cancellation to simulate a long task.
		select {
		case <-ctx.Done():
			canceled.Add(1)
			return ctx.Err()
		case <-time.After(500 * time.Millisecond):
			return nil
		}
	}

	errc := make(chan error, 1)

	var cancelOnce sync.Once
	wg := sync.WaitGroup{}
	wg.Add(len(inputs))
	for _, in := range inputs {
		in := in
		go func() {
			defer wg.Done()
			if err := doOne(in); err != nil {
				select {
				case errc <- err:
				default:
				}
				cancelOnce.Do(func() { cancel() })
			}
		}()
	}

	wg.Wait()

	wantCanceled := int64(len(inputs) - 1)
	if canceled.Load() != wantCanceled {
		return fmt.Errorf("fail-fast: expected canceled=%d, got=%d", wantCanceled, canceled.Load())
	}

	select {
	case err := <-errc:
		_ = err
		return nil
	default:
		return fmt.Errorf("fail-fast: expected at least one error")
	}
}

