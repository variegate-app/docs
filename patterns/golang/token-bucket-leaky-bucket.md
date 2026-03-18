# token bucket / leaky bucket

## [<<< ---](../index.md)

**Token bucket** и **leaky bucket** — два классических алгоритма rate limiting.

- **Token bucket**: токены накапливаются в баке с ограниченной ёмкостью; операция выполняется, когда есть токен.
- **Leaky bucket**: “вытекание” происходит равномерно; очередь задач “протекает” через систему с фиксированной скоростью.

Ниже — минимальные реализации через Go-каналы.

## Token bucket (через канал токенов)

```go
package main

import (
	"context"
	"time"
)

type TokenBucket struct {
	tokens chan struct{}
	interval time.Duration
}

func NewTokenBucket(rate time.Duration, capacity int) *TokenBucket {
	tb := &TokenBucket{
		tokens:   make(chan struct{}, capacity),
		interval: rate,
	}

	// Начальное наполнение можно настроить: здесь наполняем полностью.
	for i := 0; i < capacity; i++ {
		tb.tokens <- struct{}{}
	}

	go func() {
		t := time.NewTicker(rate)
		defer t.Stop()
		for range t.C {
			select {
			case tb.tokens <- struct{}{}:
				// Токен добавлен, если бак не полный.
			default:
				// Бак полный — токен “не доливаем”.
			}
		}
	}()

	return tb
}

func (tb *TokenBucket) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tb.tokens:
		return nil
	}
}
```

## Leaky bucket (через очередь/«протекание»)

```go
package main

import (
	"context"
	"time"
)

// LeakyBucket пропускает операции с постоянным интервалом.
// requests можно подавать сколько угодно (пока в очереди хватает места),
// но сами выполнения будут “вытекать” равномерно.
type LeakyBucket struct {
	queue chan struct{}
	interval time.Duration
}

func NewLeakyBucket(queueSize int, interval time.Duration) *LeakyBucket {
	lb := &LeakyBucket{
		queue:   make(chan struct{}, queueSize),
		interval: interval,
	}

	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for range t.C {
			select {
			case <-lb.queue:
				// «протекли» очередную операцию
			default:
				// очереди нет — ничего не делаем
			}
		}
	}()

	return lb
}

func (lb *LeakyBucket) Enqueue(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case lb.queue <- struct{}{}:
		return nil
	default:
		// Очередь заполнена — ограничиваем нагрузку.
		return context.DeadlineExceeded
	}
}
```

С leaky bucket часто проще организовать “буферизацию” запросов и контролировать,
что именно будет выполняться дальше, а token bucket — особенно удобен, когда нужно
разрешать кратковременные всплески (burst).

