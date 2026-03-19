package graceful_shutdown

import (
	stdctx "context"
	"fmt"
	"sync"
	"time"
)

func Run(ctx stdctx.Context) error {
	ctx, cancel := stdctx.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	const workers = 3
	jobs := make(chan int, 20)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case j, ok := <-jobs:
					if !ok {
						return
					}
					_ = j // simulate work
					time.Sleep(5 * time.Millisecond)
				}
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs) // stop accepting new jobs
	wg.Wait()
	fmt.Println("graceful-shutdown: workers drained and exited")
	return nil
}

