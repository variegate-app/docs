# Строитель (Builder)

## [<<< ---](../../index.md)

## Назначение

Builder позволяет пошагово собирать сложный объект, отделяя процесс построения от итогового представления.

Клиент управляет конфигурацией шагов, а Director (или логика сборки) отвечает за порядок выполнения шагов.

## Когда использовать

- когда объект сложный и его создание требует много шагов;
- когда необходимо собрать разные варианты одного и того же продукта;
- когда хочется уменьшить количество параметров в конструкторе и сделать сборку читаемой.

## Пример реализации на Go

```go
package main

import "fmt"

// Product — сложный объект.
type Product struct {
	CPU   string
	Memory string
	OS    string
}

// Builder — шаги сборки продукта.
type Builder interface {
	SetCPU(string) Builder
	SetMemory(string) Builder
	SetOS(string) Builder
	Build() Product
}

type ServerBuilder struct {
	p Product
}

func (b ServerBuilder) SetCPU(v string) Builder {
	b.p.CPU = v
	return b
}

func (b ServerBuilder) SetMemory(v string) Builder {
	b.p.Memory = v
	return b
}

func (b ServerBuilder) SetOS(v string) Builder {
	b.p.OS = v
	return b
}

func (b ServerBuilder) Build() Product { return b.p }

func main() {
	p := ServerBuilder{}.
		SetCPU("8-core").
		SetMemory("32GB").
		SetOS("linux").
		Build()

	fmt.Printf("built: %+v\n", p)
}
```

