package map_filter

import (
	stdctx "context"
	"fmt"
)

func Map(done <-chan struct{}, in <-chan int, mapper func(int) int) <-chan int {
	out := make(chan int, 16)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
				return
			case out <- mapper(v):
			}
		}
	}()
	return out
}

func Filter(done <-chan struct{}, in <-chan int, pred func(int) bool) <-chan int {
	out := make(chan int, 16)
	go func() {
		defer close(out)
		for v := range in {
			if !pred(v) {
				continue
			}
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()
	return out
}

func Run(ctx stdctx.Context) error {
	done := ctx.Done()

	// src: 1..6
	src := make(chan int, 6)
	for i := 1; i <= 6; i++ {
		src <- i
	}
	close(src)

	// evens -> square
	evens := Filter(done, src, func(v int) bool { return v%2 == 0 })
	squares := Map(done, evens, func(v int) int { return v * v })

	var got []int
	for v := range squares {
		got = append(got, v)
	}

	expected := []int{4, 16, 36}
	if len(got) != len(expected) {
		return fmt.Errorf("map-filter: expected %v got %v", expected, got)
	}
	for i := range expected {
		if got[i] != expected[i] {
			return fmt.Errorf("map-filter: mismatch at idx %d: expected=%v got=%v", i, expected, got)
		}
	}
	return nil
}

