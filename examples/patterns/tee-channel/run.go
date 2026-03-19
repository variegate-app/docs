package tee_channel

import (
	stdctx "context"
	"fmt"
	"sort"
)

func Tee(done <-chan struct{}, in <-chan int) (<-chan int, <-chan int) {
	// Buffered outputs prevent deadlocks if one consumer is slower.
	out1 := make(chan int, 16)
	out2 := make(chan int, 16)

	go func() {
		defer close(out1)
		defer close(out2)

		for v := range in {
			select {
			case <-done:
				return
			case out1 <- v:
			}
			select {
			case <-done:
				return
			case out2 <- v:
			}
		}
	}()

	return out1, out2
}

func Run(ctx stdctx.Context) error {
	done := ctx.Done()

	in := make(chan int, 8)
	for i := 0; i < 5; i++ {
		in <- i
	}
	close(in)

	out1, out2 := Tee(done, in)

	var a, b []int
	for v := range out1 {
		a = append(a, v)
	}
	for v := range out2 {
		b = append(b, v)
	}

	if len(a) != 5 || len(b) != 5 {
		return fmt.Errorf("tee-channel: expected 5 values, got out1=%d out2=%d", len(a), len(b))
	}

	sort.Ints(a)
	sort.Ints(b)
	for i := 0; i < 5; i++ {
		if a[i] != b[i] || a[i] != i {
			return fmt.Errorf("tee-channel: mismatch at idx %d: out1=%v out2=%v", i, a, b)
		}
	}
	return nil
}

