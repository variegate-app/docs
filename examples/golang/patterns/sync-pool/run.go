package sync_pool

import (
	stdctx "context"
	"bytes"
	"fmt"
	"sync"
)

func Run(ctx stdctx.Context) error {
	_ = ctx

	var pool = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}

	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	buf.WriteString("hello")
	if buf.String() != "hello" {
		return fmt.Errorf("sync-pool: expected buffer to contain hello, got %q", buf.String())
	}

	// Return with clean state.
	buf.Reset()
	pool.Put(buf)

	buf2 := pool.Get().(*bytes.Buffer)
	if buf2.Len() != 0 {
		return fmt.Errorf("sync-pool: expected empty buffer after reuse, got len=%d", buf2.Len())
	}

	// Cleanup.
	buf2.Reset()
	pool.Put(buf2)
	return nil
}

