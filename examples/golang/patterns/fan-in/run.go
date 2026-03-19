package fan_in

import (
	stdctx "context"
	"fmt"
	"sort"
	"sync"
)

func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 32)

	send := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Run(ctx stdctx.Context) error {
	// Avoid goroutine leaks on cancellation.
	done := ctx.Done()

	c1 := make(chan int, 3)
	c1 <- 1
	c1 <- 2
	c1 <- 3
	close(c1)

	c2 := make(chan int, 2)
	c2 <- 10
	c2 <- 11
	close(c2)

	c3 := make(chan int, 1)
	c3 <- -1
	close(c3)

	out := Merge(c1, c2, c3)
	var got []int
	for {
		select {
		case <-done:
			return fmt.Errorf("fan-in: canceled: %w", ctx.Err())
		case v, ok := <-out:
			if !ok {
				sort.Ints(got)
				expected := []int{-1, 1, 2, 3, 10, 11}
				sort.Ints(expected)
				if len(got) != len(expected) {
					return fmt.Errorf("fan-in: expected %d vals, got %d", len(expected), len(got))
				}
				for i := range expected {
					if got[i] != expected[i] {
						return fmt.Errorf("fan-in: mismatch at %d: got=%v expected=%v", i, got, expected)
					}
				}
				return nil
			}
			got = append(got, v)
		}
	}
}

