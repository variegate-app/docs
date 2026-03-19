package generator

import (
	stdctx "context"
	"fmt"
)

func Count(ctx stdctx.Context, start int, end int) <-chan int {
	out := make(chan int, 8)
	go func() {
		defer close(out)
		for i := start; i <= end; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()
	return out
}

func Run(ctx stdctx.Context) error {
	out := Count(ctx, 1, 5)
	var got []int
	for v := range out {
		got = append(got, v)
	}
	expected := []int{1, 2, 3, 4, 5}
	if len(got) != len(expected) {
		return fmt.Errorf("generator: expected %v, got %v", expected, got)
	}
	for i := range expected {
		if got[i] != expected[i] {
			return fmt.Errorf("generator: mismatch at %d: got=%v expected=%v", i, got, expected)
		}
	}
	return nil
}

