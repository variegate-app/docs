package or_channel

import (
	stdctx "context"
	"fmt"
	"sync"

)

func OrDone(channels ...<-chan struct{}) <-chan struct{} {
	out := make(chan struct{})

	var once sync.Once
	for _, ch := range channels {
		c := ch
		go func() {
			<-c
			once.Do(func() { close(out) })
		}()
	}

	return out
}

func Run(ctx stdctx.Context) error {
	done1 := make(chan struct{})
	done2 := make(chan struct{})

	aggDone := OrDone(done1, done2)

	// Trigger completion.
	close(done2)

	select {
	case <-aggDone:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("or-channel: aggDone not closed: %w", ctx.Err())
	}
}

