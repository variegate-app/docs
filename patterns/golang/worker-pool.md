# Worker Pool

## [<<< ---](../index.md)

Паттерн **Worker Pool** описывает набор одинаковых worker‑горутин, считывающих задачи из общего канала и выполняющих их параллельно.

Он полезен, когда:

- есть большое количество однотипных задач;
- мы хотим ограничить максимальное число параллельных исполнителей.

```go
type Task func() error

// StartWorkerPool запускает poolSize воркеров, читающих задачи из tasks.
// Ошибки, если нужны, можно отправлять в отдельный канал.
func StartWorkerPool(done <-chan struct{}, poolSize int, tasks <-chan Task) {
	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case t, ok := <-tasks:
					if !ok {
						return
					}
					_ = t() // обработку ошибки можно добавить при необходимости
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		// здесь можно закрыть канал результатов, если он есть
	}()
}
```

Такой пул можно переиспользовать для обработки запросов, задач из очереди, файлов и т.д.

