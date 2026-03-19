package errgroup

import (
	stdctx "context"

	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func Run(ctx stdctx.Context) error {
	g, egCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		time.Sleep(20 * time.Millisecond)
		return errors.New("errgroup: first error")
	})

	g.Go(func() error {
		select {
		case <-egCtx.Done():
			return egCtx.Err()
		case <-time.After(500 * time.Millisecond):
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("errgroup: first error observed:", err)
		return nil
	}
	return nil
}

