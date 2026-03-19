package ring_buffer

import (
	stdctx "context"
	"fmt"
	"sort"
)

func Run(ctx stdctx.Context) error {
	const cap = 3
	rb := make(chan int, cap)

	// Simulate producer: when buffer is full, drop the oldest value.
	for i := 1; i <= 5; i++ {
		select {
		case <-ctx.Done():
			return fmt.Errorf("ring-buffer: canceled: %w", ctx.Err())
		default:
		}

		select {
		case rb <- i:
		default:
			// drop oldest
			<-rb
			rb <- i
		}
	}

	// Drain what's left (last cap elements).
	var got []int
Drain:
	for {
		select {
		case v := <-rb:
			got = append(got, v)
		default:
			break Drain
		}
	}

	expected := []int{3, 4, 5}
	sort.Ints(got)
	sort.Ints(expected)

	if len(got) != len(expected) {
		return fmt.Errorf("ring-buffer: expected %v, got %v", expected, got)
	}
	for i := range expected {
		if got[i] != expected[i] {
			return fmt.Errorf("ring-buffer: mismatch at idx %d: expected=%v got=%v", i, expected, got)
		}
	}
	return nil
}

