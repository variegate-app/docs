package worker_pool

import (
	stdctx "context"
	"fmt"
	"sync"
	"time"
)

func Run(ctx stdctx.Context) error {
	type Task func() error

	done := make(chan struct{})
	defer close(done)

	tasks := make(chan Task)

	const poolSize = 4
	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case t, ok := <-tasks:
					if !ok {
						return
					}
					_ = t()
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		i := i
		tasks <- func() error {
			time.Sleep(2 * time.Millisecond)
			_ = i
			return nil
		}
	}
	close(tasks) // allow workers to exit
	wg.Wait()
	fmt.Println("worker-pool: executed tasks and exited")
	return nil
}

