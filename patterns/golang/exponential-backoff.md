# Exponential backoff

## [<<< ---](../index.md)

Паттерн **Exponential backoff** используется при повторных попытках (retry) неуспешных операций, чтобы:

- не перегружать внешний сервис частыми запросами;
- дать системе время «отдохнуть» и восстановиться.

Задержка между повторами растёт экспоненциально: `base * 2^n`.

```go
func RetryWithBackoff(ctx context.Context, maxRetries int, baseDelay time.Duration, fn func() error) error {
	var err error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if err = fn(); err == nil {
			return nil
		}

		// последняя попытка — выходим
		if attempt == maxRetries {
			break
		}

		// считаем задержку
		delay := baseDelay * (1 << attempt) // base * 2^attempt

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
	}

	return err
}
```

Такой подход часто применяется для HTTP‑клиентов, запросов к БД и любым нестабильным внешним зависимостям.

## Типовые ошибки / антипаттерны
- Повторять без ограничений: бесконечные retry вместо управляемого maxRetries.
- Не учитывать `ctx.Done()`: retry может продолжаться даже после отмены/таймаута.
- Делать один и тот же backoff всем клиентам синхронно (thundering herd).

## Практический чеклист
- Есть ограничение `maxRetries` и понятная политика “когда прекращаем”.
- `ctx` используется в каждой задержке между попытками.
- Для очереди/серверов учтены лимиты: retry не усиливает нагрузку.
- При необходимости добавляется jitter (рандомизация задержки).

## Как адаптировать под кейс
- Для idempotent операций retry обычно безопаснее.
- Для внешних сервисов сочетайте backoff с circuit breaker/fail-fast.

## Связанные паттерны
- `./jitter.md`
- `./fail-fast.md`
- `./cancellation.md`

