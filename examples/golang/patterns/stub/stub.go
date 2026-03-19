package stub

import (
	stdctx "context"
	"fmt"
)

// Run is a compile-time safe placeholder for patterns whose runnable examples
// are not implemented yet.
func Run(ctx stdctx.Context, name string) error {
	_ = ctx
	fmt.Printf("[stub] %s example is not implemented yet in examples/. You can still use the markdown reference.\n", name)
	return nil
}

