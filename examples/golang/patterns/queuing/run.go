package queuing

import (
	stdctx "context"
	"fmt"
	"time"
)

func Run(ctx stdctx.Context) error {
	done := ctx.Done()

	queue := make(chan int, 4)

	processed := make([]int, 0, 6)
	errc := make(chan error, 1)

	go func() {
		for {
			select {
			case <-done:
				return
			case v, ok := <-queue:
				if !ok {
					errc <- nil
					return
				}
				// Single consumer => order preserved.
				time.Sleep(1 * time.Millisecond)
				processed = append(processed, v)
			}
		}
	}()

	input := []int{1, 2, 3, 4, 5, 6}
	for _, v := range input {
		select {
		case <-done:
			return fmt.Errorf("queuing: canceled: %w", ctx.Err())
		case queue <- v:
		}
	}
	close(queue)

	select {
	case err := <-errc:
		if err != nil {
			return err
		}
	case <-done:
		return fmt.Errorf("queuing: timeout: %w", ctx.Err())
	}

	if len(processed) != len(input) {
		return fmt.Errorf("queuing: expected %d items, got %d", len(input), len(processed))
	}
	for i := range input {
		if processed[i] != input[i] {
			return fmt.Errorf("queuing: mismatch at idx %d: expected=%d got=%d", i, input[i], processed[i])
		}
	}
	return nil
}

