# Deadline

## [<<< ---](../index.md)

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

## Типовые ошибки / антипаттерны
- Задавать дедлайн, но не передавать `ctx` в нижние уровни (DB/RPC/HTTP), из-за чего отмена не работает.
- Вызывать операции, которые игнорируют `ctx.Done()` (получается зависание поверх дедлайна).
- Возвращать “успех” после дедлайна из-за неверной логики select/ошибок.

## Практический чеклист
- Используется `context.WithDeadline/WithTimeout` и вызывается `cancel()` по завершении.
- В длинных операциях есть выборка по `ctx.Done()` для корректного выхода.
- Внешние вызовы принимают `ctx` (а не создают свои таймауты “внутри” без контроля).
- Ошибки маппятся: различаете таймаут/отмену (`ctx.Err()`) и доменные ошибки.

## Как адаптировать под кейс
- Для RPC/HTTP дедлайн обычно выставляется из SLA верхнего уровня.
- Если дедлайн слишком короткий для worst-case — добавьте retry/backoff и/или шарпинг батчей.

## Связанные паттерны
- `./cancellation.md`
- `./explicit-cancellation.md`
- `./graceful-shutdown.md`
