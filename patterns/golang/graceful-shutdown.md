# graceful shutdown / drain каналов

## [<<< ---](../index.md)

**Graceful shutdown** — это корректное завершение системы:

- прекращаем принимать новые задачи;
- даём уже принятым задачам завершиться;
- отменяем долгие операции по `context`;
- дожидаемся остановки воркеров и закрываем каналы, чтобы не было утечек.

На уровне каналов часто это выглядит как связка `done/context` + `WaitGroup` + правильные `close()`.

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(ctx context.Context, jobs <-chan Job, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobs:
			if !ok {
				// jobs закрыт — больше задач нет
				return
			}
			// имитация работы
			time.Sleep(50 * time.Millisecond)
			fmt.Println("worker", id, "done job", j.ID)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	const workers = 4
	jobs := make(chan Job, 20)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(ctx, jobs, i, &wg)
	}

	// Подаём задачи (пока система «жива»).
	for i := 0; i < 50; i++ {
		jobs <- Job{ID: i}
	}

	// “Остановка принятия новых задач”: закрываем jobs.
	close(jobs)

	// Wait: дождаться окончания работы воркеров (или остановки по ctx.Done()).
	wg.Wait()
}
```

Если нужно “drain” результатов: отдельный consumer обязан дочитать оставшиеся значения из результата‑канала до его `close()`.

