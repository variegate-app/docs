# Context

## [<<< ---](../gochan.md)

Паттерн **Context** в Go реализован стандартным пакетом `context` и используется для:

- распространения сигналов отмены;
- задания дедлайнов и таймаутов;
- передачи сквозных метаданных (request id, пользователь и т.п.).

Контекст обычно передаётся первым аргументом в функции, выполняющие работу.

```go
func doWork(ctx context.Context, id int) error {
	select {
	case <-time.After(200 * time.Millisecond):
		// работа завершилась вовремя
		return nil
	case <-ctx.Done():
		// контекст отменён или сработал дедлайн
		return ctx.Err()
	}
}

func main() {
	// создаём контекст с таймаутом 100 мс
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := doWork(ctx, 1); err != nil {
		fmt.Println("work cancelled:", err)
	}
}
```

Все функции, работающие с внешними ресурсами (БД, сеть, RPC), должны принимать `context.Context`, чтобы корректно реагировать на отмену и дедлайны.

