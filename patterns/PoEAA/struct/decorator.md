# Декоратор (Decorator)

## [<<< ---](../../index.md)

## Назначение

Декоратор динамически добавляет объекту новые обязанности, не меняя его исходный код и позволяя комбинировать поведение.

По сути, вы «оборачиваете» объект в слой (декоратор), который перехватывает вызовы и добавляет функциональность до/после делегирования исходному объекту.

## Когда использовать Decorator

- когда нужно расширять поведение объекта «поверх» базовой реализации (композиция вместо наследования);
- когда комбинации обязанностей должны собираться гибко во время выполнения;
- когда нельзя (или не хочется) модифицировать исходный класс.

## Пример реализации на Go

```go
package main

import "fmt"

// Component — базовое поведение.
type Component interface {
	Operation() string
}

type Base struct{}

func (b Base) Operation() string { return "base" }

// Decorator — “обёртка” над Component.
type Decorator struct {
	c Component
}

func (d Decorator) Operation() string { return d.c.Operation() }

// LoggingDecorator добавляет поведение поверх.
type LoggingDecorator struct {
	Decorator
}

func (l LoggingDecorator) Operation() string {
	fmt.Println("log: before")
	out := l.c.Operation()
	fmt.Println("log: after")
	return out + " + logged"
}

func main() {
	var c Component = Base{}
	c = LoggingDecorator{Decorator{c: c}}
	fmt.Println(c.Operation())
}
```
