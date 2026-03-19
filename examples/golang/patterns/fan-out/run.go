package fan_out

import (
	stdctx "context"
	"fmt"
)

func Split(in <-chan int, n int) []<-chan int {
	if n <= 0 {
		panic("n must be > 0")
	}

	outs := make([]chan int, n)
	for i := 0; i < n; i++ {
		outs[i] = make(chan int, 16)
	}

	go func() {
		defer func() {
			for _, ch := range outs {
				close(ch)
			}
		}()

		i := 0
		for v := range in {
			outs[i%n] <- v
			i++
		}
	}()

	res := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		res[i] = outs[i]
	}
	return res
}

func Run(ctx stdctx.Context) error {
	in := make(chan int, 16)
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)

	outs := Split(in, 3)

	seen := make([]bool, 10)
	count := 0
	for _, ch := range outs {
		for v := range ch {
			if v < 0 || v >= 10 {
				return fmt.Errorf("fan-out: unexpected value %d", v)
			}
			if seen[v] {
				return fmt.Errorf("fan-out: duplicated value %d", v)
			}
			seen[v] = true
			count++
		}
	}

	for i := 0; i < 10; i++ {
		if !seen[i] {
			return fmt.Errorf("fan-out: missing value %d", i)
		}
	}
	if count != 10 {
		return fmt.Errorf("fan-out: expected count=10, got %d", count)
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("fan-out: canceled: %w", ctx.Err())
	default:
	}
	return nil
}

