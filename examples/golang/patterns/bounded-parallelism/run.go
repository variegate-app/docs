package bounded_parallelism

import (
	stdctx "context"
	"fmt"
	"sync"
)

func Run(ctx stdctx.Context) error {
	done := make(chan struct{})
	defer close(done)

	paths := make(chan int)
	results := make(chan int)

	const workers = 4
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range paths {
				select {
				case <-done:
					return
				case results <- v * v:
				}
			}
		}()
	}

	go func() {
		defer close(results)
		wg.Wait()
	}()

	go func() {
		defer close(paths)
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				return
			case paths <- i:
			}
		}
	}()

	total := 0
	for r := range results {
		total += r
	}
	fmt.Println("bounded-parallelism: total =", total)
	return nil
}

