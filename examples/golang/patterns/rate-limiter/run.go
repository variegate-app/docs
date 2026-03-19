package rate_limiter

import (
	stdctx "context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func Run(ctx stdctx.Context) error {
	lim := rate.NewLimiter(rate.Every(40*time.Millisecond), 2)
	for i := 0; i < 3; i++ {
		if err := lim.Wait(ctx); err != nil {
			return err
		}
		fmt.Println("rate-limiter: allow", i)
	}
	return nil
}

