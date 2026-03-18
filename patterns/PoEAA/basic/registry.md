# Registry (Реестр)

## [<<< ---](../../index.md)

Хорошо известный объект, который используется другими объектами для получения общих объектов и сервисов.

Когда нужно найти какой-нибудь объект, обычно начинают с другого объекта, связанного с целевым. Например, если нужно найти все счета для покупателя, начинают, как раз с покупателя и используют его метод получения счетов. Тем не менее, в некоторых случаях нет подходящего объекта, с которого начать. Например, известен ID покупателя, но нет ссылки на него. Тогда нужен своего рода объект-поисковик, но тогда возникает вопрос - как вы найдёте сам поисковик?

Реестр (Registry) — это глобальный объект по сути своей или, по крайней мере, так выглядит - он может функционировать только будучи глобальным.

### Пример реализации на Go (Registry)

```go
package main

import (
	"fmt"
	"sync"
)

// Registry хранит “общие” сервисы по ключу.
type Registry struct {
	mu       sync.RWMutex
	services map[string]any
}

func NewRegistry() *Registry {
	return &Registry{services: map[string]any{}}
}

func (r *Registry) Register(name string, svc any) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.services[name] = svc
}

func (r *Registry) Resolve[T any](name string) (T, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.services[name]
	if !ok {
		var zero T
		return zero, false
	}
	res, ok := v.(T)
	return res, ok
}

// Пример сервисов.
type Logger struct{}

func (l Logger) Print(msg string) { fmt.Println("log:", msg) }

func main() {
	reg := NewRegistry()
	reg.Register("logger", Logger{})

	logger, ok := reg.Resolve[Logger]("logger")
	if !ok {
		return
	}

	logger.Print("hello")
}
```