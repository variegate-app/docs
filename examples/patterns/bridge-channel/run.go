package bridge_channel

import (
	stdctx "context"
	"fmt"
	"sort"
)

func Bridge(done <-chan struct{}, chanStream <-chan <-chan int) <-chan int {
	out := make(chan int, 32)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case stream, ok := <-chanStream:
				if !ok {
					return
				}
				for v := range stream {
					select {
					case <-done:
						return
					case out <- v:
					}
				}
			}
		}
	}()

	return out
}

func Run(ctx stdctx.Context) error {
	done := ctx.Done()

	chanStream := make(chan (<-chan int), 4)

	inner1 := make(chan int, 3)
	inner1 <- 1
	inner1 <- 2
	inner1 <- 3
	close(inner1)

	inner2 := make(chan int, 2)
	inner2 <- 10
	inner2 <- 11
	close(inner2)

	chanStream <- inner1
	chanStream <- inner2
	close(chanStream)

	out := Bridge(done, chanStream)
	var got []int
	for v := range out {
		got = append(got, v)
	}

	expected := []int{1, 2, 3, 10, 11}
	sort.Ints(got)
	sort.Ints(expected)

	if len(got) != len(expected) {
		return fmt.Errorf("bridge-channel: expected %d values, got %d", len(expected), len(got))
	}
	for i := range expected {
		if got[i] != expected[i] {
			return fmt.Errorf("bridge-channel: mismatch at %d: got=%v expected=%v", i, got, expected)
		}
	}
	return nil
}

