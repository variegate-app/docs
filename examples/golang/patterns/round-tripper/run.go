package round_tripper

import (
	stdctx "context"
	"fmt"
)

type Request struct {
	ID      int
	Payload string
	RespCh  chan Response
}

type Response struct {
	ID      int
	Payload string
	Err     error
}

func Run(ctx stdctx.Context) error {
	reqs := make(chan Request)

	// Server.
	serverDone := make(chan struct{})
	go func() {
		defer close(serverDone)
		for req := range reqs {
			// Simulate processing and respond on per-request channel.
			req.RespCh <- Response{
				ID:      req.ID,
				Payload: "OK: " + req.Payload,
				Err:     nil,
			}
		}
	}()

	respCh := make(chan Response, 1)
	req := Request{
		ID:       42,
		Payload: "hello",
		RespCh:   respCh,
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("round-tripper: canceled: %w", ctx.Err())
	case reqs <- req:
	}

	var resp Response
	select {
	case <-ctx.Done():
		return fmt.Errorf("round-tripper: canceled waiting response: %w", ctx.Err())
	case resp = <-respCh:
	}

	if resp.Err != nil {
		return fmt.Errorf("round-tripper: unexpected error: %v", resp.Err)
	}
	if resp.ID != req.ID {
		return fmt.Errorf("round-tripper: expected ID=%d got=%d", req.ID, resp.ID)
	}
	if resp.Payload != "OK: "+req.Payload {
		return fmt.Errorf("round-tripper: payload mismatch: got=%q expected=%q", resp.Payload, "OK: "+req.Payload)
	}

	close(reqs)
	<-serverDone
	return nil
}

