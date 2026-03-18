# jitter для retry/backoff

## [<<< ---](../index.md)

**Jitter** — это добавление случайности к задержкам при retry/backoff.

Зачем: когда много горутин одновременно делают повторные попытки, одинаковый backoff может привести к эффекту **thundering herd** (все “просыпаются” в одно и то же время). Случайность разносит повторные попытки по времени.

Пример: «полный jitter» (full jitter) и «равномерный jitter» (equal jitter).

```go
package main

import (
	"math/rand"
	"time"
)

func FullJitter(base time.Duration) time.Duration {
	// rand[0..base)
	return time.Duration(rand.Int63n(int64(base)))
}

func EqualJitter(min, max time.Duration) time.Duration {
	if max <= min {
		return min
	}
	half := (max - min) / 2
	if half <= 0 {
		return min
	}
	// min + rand[0..half)
	return min + time.Duration(rand.Int63n(int64(half)))
}

// BackoffStep возвращает следующую задержку для retry:
// например, можно сделать base на основе попытки и добавить равномерный jitter.
func BackoffStep(attempt int, base, max time.Duration) time.Duration {
	// Упрощённая экспонента: base * 2^attempt (с ограничением max).
	// Затем добавляем равномерный jitter в диапазон [base, max].
	// (В реальном коде часто комбинируют с exponential backoff.)
	delay := base * (1 << attempt)
	if delay > max {
		delay = max
	}
	return EqualJitter(delay/2, delay)
}
```

