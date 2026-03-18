# Наблюдатель (Observer)

## [<<< ---](../../index.md)

## Назначение

Observer связывает объект-издатель (subject) и множество подписчиков (обработчиков), чтобы подписчики автоматически реагировали на события.

Издатель не знает конкретных классов подписчиков.

## Когда использовать

- когда нужно организовать реакцию на события без жёсткой связанности;
- когда подписчиков может быть много или они появляются динамически;
- когда важно, чтобы события рассылались независимо от того, какие конкретно обработчики стоят.

## Пример реализации на Go

```go
package main

import "fmt"

type Event struct {
	Topic string
	Data  any
}

type Observer interface {
	Update(e Event)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(e Event) {
	for _, o := range s.observers {
		o.Update(e)
	}
}

type Logger struct{}

func (l Logger) Update(e Event) {
	fmt.Println("log:", e.Topic, e.Data)
}

type Counter struct{ n int }

func (c *Counter) Update(e Event) {
	c.n++
	fmt.Println("count:", c.n)
}

func main() {
	sub := &Subject{}
	sub.Attach(Logger{})
	sub.Attach(&Counter{})

	sub.Notify(Event{Topic: "user.created", Data: "u1"})
	sub.Notify(Event{Topic: "user.created", Data: "u2"})
}
```

