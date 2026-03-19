# debounce

## [<<< ---](../index.md)

`debounce` — паттерн “схлопывания” событий: пока события продолжают поступать, таймер постоянно сбрасывается; реальное действие выполняется только после паузы длиной `delay`.

```go
package main

import (
	"sync"
	"time"
)

type Debouncer struct {
	delay time.Duration

	mu     sync.Mutex
	timer  *time.Timer
	closed bool

	fn func()
}

func NewDebouncer(delay time.Duration, fn func()) *Debouncer {
	return &Debouncer{delay: delay, fn: fn}
}

// Call планирует выполнение fn после задержки.
// Если Call вызывается снова до срабатывания таймера — таймер перезапускается.
func (d *Debouncer) Call() {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.closed {
		return
	}

	if d.timer == nil {
		d.timer = time.NewTimer(d.delay)
		go d.wait()
		return
	}

	if !d.timer.Stop() {
		// Если таймер уже почти сработал — осушаем канал, чтобы не было гонок.
		select {
		case <-d.timer.C:
		default:
		}
	}
	d.timer.Reset(d.delay)
}

func (d *Debouncer) wait() {
	d.mu.Lock()
	t := d.timer
	d.mu.Unlock()

	<-t.C
	d.mu.Lock()
	if d.closed {
		d.mu.Unlock()
		return
	}
	d.mu.Unlock()

	d.fn()
}

func (d *Debouncer) Stop() {
	d.mu.Lock()
	d.closed = true
	if d.timer != nil {
		d.timer.Stop()
	}
	d.mu.Unlock()
}

func main() {
	// пример: “выполнять пересчет только когда пользователь перестал печатать 300 мс”
	deb := NewDebouncer(300*time.Millisecond, func() {
		// делать работу
	})

	deb.Call()
	// ... новые события сбрасывают таймер ...
	// deb.Stop() при необходимости
}
```

## Типовые ошибки / антипаттерны
- Забывать про `Stop()`: таймер может “догнать” действие после shutdown/удаления объекта.
- Держать mutex слишком долго: `Call()`/`wait()` конфликтуют, растёт латентность.
- Не защищать повторные вызовы `Call()` от гонок (в примере это решено `mu` и флагом `closed`).

## Практический чеклист
- Есть явный метод остановки (`Stop`) при жизненном цикле объекта.
- Поведение определено: при частых событиях действие выполняется один раз после “тишины” `delay`.
- Таймер корректно reset-ится (осторожно с `timer.Stop()` и чтением из `timer.C`).
- Внутри `fn()` нет долгих блокировок, иначе debounce потеряет смысл.

## Как адаптировать под кейс
- Для ввода пользователя (typing/search) debounce обычно улучшает UX и снижает нагрузку.
- Если нужен “лимит сверху” (не реже N и не чаще M) — combine debounce с rate limiting/throttle.

## Связанные паттерны
- `./rate-limiter.md`
- `./cancellation.md`

