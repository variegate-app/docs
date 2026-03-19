package debounce

import (
	stdctx "context"
	"fmt"
	"sync/atomic"
	"time"
)

type Debouncer struct {
	delay time.Duration

	// protects timer reset
	resetCh chan struct{}
	stopCh  chan struct{}
	fn      func()
}

func NewDebouncer(delay time.Duration, fn func()) *Debouncer {
	return &Debouncer{
		delay:   delay,
		resetCh: make(chan struct{}, 1),
		stopCh:  make(chan struct{}),
		fn:      fn,
	}
}

func (d *Debouncer) Call() {
	select {
	case d.resetCh <- struct{}{}:
	default:
		// already reset scheduled
	}
}

func (d *Debouncer) Stop() {
	close(d.stopCh)
}

func (d *Debouncer) Run(ctx stdctx.Context) error {
	t := time.NewTimer(d.delay)
	if !t.Stop() {
		<-t.C
	}

	for {
		select {
		case <-ctx.Done():
			t.Stop()
			return ctx.Err()
		case <-d.stopCh:
			t.Stop()
			return nil
		case <-d.resetCh:
			// reset timer (drain if needed)
			if !t.Stop() {
				select {
				case <-t.C:
				default:
				}
			}
			t.Reset(d.delay)
		case <-t.C:
			d.fn()
		}
	}
}

func Run(ctx stdctx.Context) error {
	var calls atomic.Int64

	deb := NewDebouncer(30*time.Millisecond, func() {
		calls.Add(1)
	})

	runErrc := make(chan error, 1)
	go func() { runErrc <- deb.Run(ctx) }()

	// Burst 1: several calls close together => single execution.
	deb.Call()
	time.Sleep(10 * time.Millisecond)
	deb.Call()
	time.Sleep(10 * time.Millisecond)
	deb.Call()

	time.Sleep(70 * time.Millisecond) // > delay
	if got := calls.Load(); got != 1 {
		return fmt.Errorf("debounce: expected 1 call after burst1, got %d", got)
	}

	// Burst 2: one call => another execution.
	deb.Call()
	time.Sleep(70 * time.Millisecond)
	if got := calls.Load(); got != 2 {
		return fmt.Errorf("debounce: expected 2 calls after burst2, got %d", got)
	}

	// Stop background loop.
	deb.Stop()
	if err := <-runErrc; err != nil {
		return fmt.Errorf("debounce: runner ended with error: %v", err)
	}
	return nil
}

