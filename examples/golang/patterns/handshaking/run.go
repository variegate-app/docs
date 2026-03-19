package handshaking

import (
	stdctx "context"
	"fmt"
)

type Message struct {
	Payload string
	Ack     chan struct{}
}

func Run(ctx stdctx.Context) error {
	in := make(chan Message, 1)
	ack := make(chan struct{})

	// Receiver.
	recvDone := make(chan struct{})
	go func() {
		defer close(recvDone)
		select {
		case <-ctx.Done():
			return
		case msg := <-in:
			if msg.Payload != "hello" {
				return
			}
			close(msg.Ack)
		}
	}()

	in <- Message{Payload: "hello", Ack: ack}

	select {
	case <-ctx.Done():
		return fmt.Errorf("handshaking: canceled: %w", ctx.Err())
	case <-ack:
		// ok
	}

	select {
	case <-recvDone:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("handshaking: receiver did not finish: %w", ctx.Err())
	}
}

