package token_bucket_leaky_bucket

import (
	stdctx "context"
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens chan struct{}
}

func NewTokenBucket(interval time.Duration, capacity int) *TokenBucket {
	tb := &TokenBucket{tokens: make(chan struct{}, capacity)}
	// fill with burst capacity
	for i := 0; i < capacity; i++ {
		tb.tokens <- struct{}{}
	}

	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for range t.C {
			select {
			case tb.tokens <- struct{}{}:
			default:
			}
		}
	}()

	return tb
}

func (tb *TokenBucket) Wait(ctx stdctx.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tb.tokens:
		return nil
	}
}

// Leaky bucket: a queue of waiting tokens that are released at fixed interval.
type LeakyBucket struct {
	sem chan struct{}
}

func NewLeakyBucket(releaseInterval time.Duration, queueSize int) *LeakyBucket {
	lb := &LeakyBucket{sem: make(chan struct{}, queueSize)}
	for i := 0; i < queueSize; i++ {
		// drain initial capacity (so queue is empty)
		// We'll release permits on ticker.
	}
	go func() {
		t := time.NewTicker(releaseInterval)
		defer t.Stop()
		for range t.C {
			select {
			case lb.sem <- struct{}{}:
			default:
			}
		}
	}()
	return lb
}

func (lb *LeakyBucket) Enqueue(ctx stdctx.Context) error {
	// Wait for permit to "leak" (execution starts on permit).
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-lb.sem:
		return nil
	}
}

func Run(ctx stdctx.Context) error {
	// Token bucket check
	ctx, cancel := stdctx.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	interval := 50 * time.Millisecond
	capacity := 2
	tb := NewTokenBucket(interval, capacity)

	start := time.Now()
	var times [4]time.Duration
	for i := 0; i < 4; i++ {
		if err := tb.Wait(ctx); err != nil {
			return fmt.Errorf("token-bucket: Wait failed: %w", err)
		}
		times[i] = time.Since(start)
	}

	// First two should be immediate (initially pre-filled).
	if times[0] > 20*time.Millisecond || times[1] > 20*time.Millisecond {
		return fmt.Errorf("token-bucket: expected burst permits quickly, got times=%v", times)
	}
	// Third permit requires at least one refill (~interval).
	if times[2] < interval-10*time.Millisecond {
		return fmt.Errorf("token-bucket: expected >= interval refill for third permit, got time[2]=%v interval=%v", times[2], interval)
	}

	// Leaky bucket check: permits come out at fixed interval.
	releaseInterval := 30 * time.Millisecond
	lb := NewLeakyBucket(releaseInterval, 2)

	// We'll enqueue 3 operations and record completion times.
	var mu sync.Mutex
	var got []time.Duration

	doOp := func() error {
		if err := lb.Enqueue(ctx); err != nil {
			return err
		}
		mu.Lock()
		got = append(got, time.Since(start))
		mu.Unlock()
		return nil
	}

	errc := make(chan error, 3)
	for i := 0; i < 3; i++ {
		go func() { errc <- doOp() }()
	}
	for i := 0; i < 3; i++ {
		if err := <-errc; err != nil {
			return fmt.Errorf("leaky-bucket: op failed: %w", err)
		}
	}

	// We expect leaked permits spaced roughly by releaseInterval.
	// Sort observed times and check consecutive deltas.
	// Because scheduling can vary, use a coarse bound.
	sortDur := func() {
		for i := 0; i < len(got); i++ {
			for j := i + 1; j < len(got); j++ {
				if got[j] < got[i] {
					got[i], got[j] = got[j], got[i]
				}
			}
		}
	}
	sortDur()
	if len(got) != 3 {
		return fmt.Errorf("leaky-bucket: expected 3 ops, got %d", len(got))
	}
	d1 := got[1] - got[0]
	d2 := got[2] - got[1]
	if d1 < releaseInterval/2 || d2 < releaseInterval/2 {
		return fmt.Errorf("leaky-bucket: expected spaced permits ~%v, deltas=%v,%v", releaseInterval, d1, d2)
	}

	return nil
}

