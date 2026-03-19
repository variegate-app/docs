package pub_sub

import (
	stdctx "context"
	"fmt"
)

func Run(ctx stdctx.Context) error {
	events := []int{1, 2, 3, 4, 5}
	bufferSize := len(events)

	sub1 := make(chan int, bufferSize)
	sub2 := make(chan int, bufferSize)

	subs := []chan<- int{sub1, sub2}

	for _, e := range events {
		for _, s := range subs {
			select {
			case <-ctx.Done():
				return fmt.Errorf("pub-sub: canceled: %w", ctx.Err())
			case s <- e:
			}
		}
	}

	close(sub1)
	close(sub2)

	readAll := func(ch <-chan int) ([]int, error) {
		var out []int
		for v := range ch {
			out = append(out, v)
		}
		return out, nil
	}

	a, _ := readAll(sub1)
	b, _ := readAll(sub2)

	if len(a) != len(events) || len(b) != len(events) {
		return fmt.Errorf("pub-sub: expected %d events, got sub1=%d sub2=%d", len(events), len(a), len(b))
	}
	for i := range events {
		if a[i] != events[i] {
			return fmt.Errorf("pub-sub: sub1 mismatch at %d: got %d expected %d", i, a[i], events[i])
		}
		if b[i] != events[i] {
			return fmt.Errorf("pub-sub: sub2 mismatch at %d: got %d expected %d", i, b[i], events[i])
		}
	}
	return nil
}

