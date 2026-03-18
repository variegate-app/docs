# Deadline

## [<<< ---](../gochan.md)

Паттерн **Deadline** задаёт жёсткий момент времени, после которого результат операции больше не нужен.

В Go это реализуется через `context.WithDeadline` или `context.WithTimeout`.

```go
func fetchWithDeadline(parent context.Context, d time.Duration) error {
	ctx, cancel := context.WithTimeout(parent, d)
	defer cancel()

	done := make(chan struct{})

	go func() {
		defer close(done)
		// имитация долгой операции
		time.Sleep(2 * time.Second)
	}()

	select {
	case <-ctx.Done():
		// дедлайн сработал раньше, чем завершилась работа
		return ctx.Err()
	case <-done:
		// операция успела завершиться до дедлайна
		return nil
	}
}
```

Deadline особенно важен для RPC/HTTP‑запросов, где клиент не должен ждать ответа бесконечно долго.

