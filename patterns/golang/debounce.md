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

