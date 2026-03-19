package singleflight

import (
	stdctx "context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func Run(ctx stdctx.Context) error {
	var sf singleflight.Group
	load := func(key string) (any, error) {
		time.Sleep(30 * time.Millisecond)
		return "value for " + key, nil
	}

	key := "user:42"
	const n = 5

	var wg sync.WaitGroup
	wg.Add(n)

	results := make([]string, n)
	shared := make([]bool, n)

	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			v, err, s := sf.Do(key, func() (any, error) { return load(key) })
			if err != nil {
				results[i] = "ERR"
				return
			}
			results[i] = v.(string)
			shared[i] = s
		}()
	}

	wg.Wait()

	fmt.Println("singleflight: results[0] =", results[0], "shared flags =", shared)
	return nil
}

