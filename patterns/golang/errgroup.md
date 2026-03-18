# errgroup

## [<<< ---](../index.md)

`errgroup` — это способ запускать несколько конкурентных задач и:

- собрать первую/единственную ошибку;
- отменять остальные задачи при первой ошибке (через общий `context`).

Обычно используется вместо “ручного” `WaitGroup` + канал ошибок.

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		return errors.New("first task failed")
	})

	g.Go(func() error {
		// Вторая задача обязана слушать ctx.Done(), чтобы корректно прерваться.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
			return nil
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("group error:", err)
	}
}
```

Ключевая идея: “первая ошибка” переводит `context` в отменённое состояние, а задачи, которые его слушают, должны остановиться.

