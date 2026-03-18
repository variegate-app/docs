# Абстрактная фабрика (Abstract Factory)

## [<<< ---](../../index.md)

## Назначение

Abstract Factory создаёт семейства связанных объектов, не указывая конкретные классы.

Она полезна, когда нужно работать с различными конфигурациями/вариантами продуктов (варианты реализации) и гарантировать согласованность между ними.

## Когда использовать

- когда в системе существуют семейства продуктов, которые должны создаваться вместе;
- когда клиенту не нужно (и/или нельзя) знать конкретные классы создаваемых объектов;
- когда варианты продукта могут меняться независимо от бизнес-логики.

## Пример реализации на Go

```go
package main

import "fmt"

// Abstract products
type Button interface{ Render() }
type Checkbox interface{ Render() }

// Factories
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Windows products
type WinButton struct{}

func (b WinButton) Render() { fmt.Println("render Windows button") }

type WinCheckbox struct{}

func (c WinCheckbox) Render() { fmt.Println("render Windows checkbox") }

type WindowsFactory struct{}

func (f WindowsFactory) CreateButton() Button     { return WinButton{} }
func (f WindowsFactory) CreateCheckbox() Checkbox { return WinCheckbox{} }

// Mac products
type MacButton struct{}

func (b MacButton) Render() { fmt.Println("render Mac button") }

type MacCheckbox struct{}

func (c MacCheckbox) Render() { fmt.Println("render Mac checkbox") }

type MacFactory struct{}

func (f MacFactory) CreateButton() Button     { return MacButton{} }
func (f MacFactory) CreateCheckbox() Checkbox { return MacCheckbox{} }

func main() {
	var factory GUIFactory = WindowsFactory{}
	factory.CreateButton().Render()
	factory.CreateCheckbox().Render()
}
```

