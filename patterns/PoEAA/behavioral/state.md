# Состояние (State)

## [<<< ---](../../index.md)

## Назначение

State меняет поведение объекта в зависимости от его внутреннего состояния.

Вместо больших условных операторов объект делегирует поведение отдельным классам состояний.

## Когда использовать

- когда поведение объекта сильно зависит от состояния и меняется со временем;
- когда логика постоянно разрастается и её трудно сопровождать;
- когда состояния можно выделить как отдельные независимые модули.

## Пример реализации на Go

```go
package main

import "fmt"

type State interface {
	Handle(*Context)
}

type Context struct {
	state State
}

func (c *Context) SetState(s State) { c.state = s }

func (c *Context) Request() {
	c.state.Handle(c)
}

type Draft struct{}

func (s Draft) Handle(c *Context) {
	fmt.Println("draft -> review")
	c.SetState(Review{})
}

type Review struct{}

func (s Review) Handle(c *Context) {
	fmt.Println("review -> published")
	c.SetState(Published{})
}

type Published struct{}

func (s Published) Handle(c *Context) {
	fmt.Println("published: no more transitions")
}

func main() {
	c := &Context{state: Draft{}}
	c.Request()
	c.Request()
	c.Request()
}
```

