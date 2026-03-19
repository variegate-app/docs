package pipeline

import (
	stdctx "context"
	"fmt"
	"sync"
)

func Run(ctx stdctx.Context) error {
	done := make(chan struct{})
	defer close(done)

	in := make(chan int)
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)
		for v := range in {
			select {
			case <-done:
				return
			case out <- v * 2:
			}
		}
	}()

	go func() {
		defer close(in)
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				return
			case in <- i:
			}
		}
	}()

	sum := 0
	for v := range out {
		sum += v
	}
	wg.Wait()
	fmt.Println("pipeline: done, sum =", sum)
	return nil
}

