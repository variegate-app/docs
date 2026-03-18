# Фабричный метод (Factory Method)

## [<<< ---](../../index.md)

## Назначение

Factory Method делегирует подклассам решение о создании объекта, сохраняя при этом единый интерфейс создания.

Базовый код описывает «когда» создавать, а фабричный метод описывает «как именно».

## Когда использовать

- когда класс заранее не знает, какие именно объекты ему нужны;
- когда нужно вынести вариативную часть создания в расширения (подклассы);
- когда создание объектов должно быть расширяемым без изменения основного алгоритма.

## Пример реализации на Go

```go
package main

import "fmt"

type Product interface {
	Name() string
}

type ProductA struct{}

func (ProductA) Name() string { return "ProductA" }

type ProductB struct{}

func (ProductB) Name() string { return "ProductB" }

// Creator описывает алгоритм и делегирует создание Product методу.
type Creator interface {
	FactoryMethod() Product
	DoWork()
}

type BaseCreator struct {
	impl Creator
}

func (b BaseCreator) DoWork() {
	p := b.impl.FactoryMethod()
	fmt.Println("work with:", p.Name())
}

type CreatorA struct{}

func (CreatorA) FactoryMethod() Product { return ProductA{} }

type CreatorB struct{}

func (CreatorB) FactoryMethod() Product { return ProductB{} }

func main() {
	a := CreatorA{}
	b := CreatorB{}

	BaseCreator{impl: a}.DoWork()
	BaseCreator{impl: b}.DoWork()
}
```

