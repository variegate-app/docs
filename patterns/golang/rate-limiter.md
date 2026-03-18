# throttle / rate limiter

## [<<< ---](../index.md)

Паттерн **rate limiter** ограничивает частоту выполнения операции (например, запросов к API), чтобы:

- защититься от всплесков нагрузки;
- уважать лимиты внешнего сервиса;
- стабилизировать систему.

Один из наиболее распространённых вариантов — токен-бакет (token bucket). В Go удобно использовать готовую реализацию из `golang.org/x/time/rate`.

```go
package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// max=10 действий сразу (burst), replenishment=1 токен в 100ms
	lim := rate.NewLimiter(rate.Every(100*time.Millisecond), 10)

	ctx := context.Background()
	for i := 0; i < 30; i++ {
		if err := lim.Wait(ctx); err != nil {
			panic(err)
		}
		fmt.Println("allow", i, time.Now().Format(time.StampMilli))
	}
}
```

Если нужно “разрешать только когда есть токен” без ожидания — используй `lim.Allow()`.

