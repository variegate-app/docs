package filter

import (
	stdctx "context"
	"fmt"
)

func Filter(done <-chan struct{}, in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int, 16)
	go func() {
		defer close(out)
		for v := range in {
			if !pred(v) {
				continue
			}
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()
	return out
}

func Run(ctx stdctx.Context) error {
	done := ctx.Done()

	in := make(chan int, 8)
	// Includes negatives and positives.
	for _, v := range []int{-2, -1, 0, 1, 2, 3} {
		in <- v
	}
	close(in)

	out := Filter(done, in, func(v int) bool { return v > 0 })

	var got []int
	for v := range out {
		got = append(got, v)
	}

	expected := []int{1, 2, 3}
	if len(got) != len(expected) {
		return fmt.Errorf("filter: expected %v, got %v", expected, got)
	}
	for i := range expected {
		if got[i] != expected[i] {
			return fmt.Errorf("filter: mismatch at %d: expected=%v got=%v", i, expected, got)
		}
	}
	return nil
}

