# Посредник (Mediator)

## [<<< ---](../../index.md)

## Назначение

Посредник централизует взаимодействие между компонентами, уменьшая связанность (coupling).

Компоненты общаются через посредника вместо прямых ссылок друг на друга.

## Когда использовать Mediator

- когда между компонентами много взаимосвязей и сложно поддерживать их напрямую;
- когда требуется централизовать логику маршрутизации событий/команд;
- когда упрощение зависимостей важнее локальной оптимизации.

## Пример реализации на Go

```go
package main

import "fmt"

type Message struct {
	From string
	Text string
}

type Colleague interface {
	Name() string
	Send(text string)
	Receive(msg Message)
}

// Mediator маршрутизирует сообщения между коллегами.
type Mediator interface {
	Dispatch(msg Message)
}

type SimpleMediator struct {
	clients map[string]Colleague
}

func (m *SimpleMediator) Dispatch(msg Message) {
	for name, c := range m.clients {
		if name != msg.From {
			c.Receive(msg)
		}
	}
}

type Client struct {
	name string
	med  Mediator
}

func (c Client) Name() string { return c.name }

func (c Client) Send(text string) {
	c.med.Dispatch(Message{From: c.name, Text: text})
}

func (c Client) Receive(msg Message) {
	fmt.Printf("%s got from %s: %s\n", c.name, msg.From, msg.Text)
}

func main() {
	med := &SimpleMediator{clients: map[string]Colleague{}}
	a := Client{name: "A", med: med}
	b := Client{name: "B", med: med}
	med.clients[a.Name()] = a
	med.clients[b.Name()] = b

	a.Send("hello")
}
```

