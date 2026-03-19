package take_first_n

import (
	stdctx "context"
	"fmt"
)

func Run(ctx stdctx.Context) error {
	done := make(chan struct{})
	defer close(done)

	// Generator: 0..9
	src := make(chan int)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				return
			case src <- i:
			}
		}
	}()

	takeInt := func(done <-chan struct{}, valueStream <-chan int, n int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 0; i < n; i++ {
				select {
				case <-done:
					return
				case v, ok := <-valueStream:
					if !ok {
						return
					}
					out <- v
				}
			}
		}()
		return out
	}

	out := takeInt(done, src, 5)
	sum := 0
	for v := range out {
		sum += v
	}
	fmt.Println("take-first-n: sum =", sum)
	return nil
}

