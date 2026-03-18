# Итератор (Iterator)

## [<<< ---](../../index.md)

## Назначение

Итератор предоставляет единый способ последовательно обходить элементы агрегата, не раскрывая его внутреннюю структуру.

Это упрощает работу с коллекциями и делает клиентский код более независимым от реализации.

## Когда использовать

- когда нужно скрыть структуру коллекции и предоставить единый API обхода;
- когда коллекция может менять способ хранения, не затрагивая клиентов;
- когда требуется несколько независимых обходов одновременно.

## Пример реализации на Go

```go
package main

import "fmt"

// Iterator — минимальный интерфейс итератора.
type Iterator[T any] interface {
	Next() bool
	Value() T
}

type SliceIterator[T any] struct {
	s   []T
	i   int
	val T
}

func NewSliceIterator[T any](s []T) *SliceIterator[T] {
	return &SliceIterator[T]{s: s, i: -1}
}

func (it *SliceIterator[T]) Next() bool {
	it.i++
	if it.i >= len(it.s) {
		return false
	}
	it.val = it.s[it.i]
	return true
}

func (it *SliceIterator[T]) Value() T { return it.val }

func main() {
	it := NewSliceIterator([]string{"a", "b", "c"})
	for it.Next() {
		fmt.Println(it.Value())
	}
}
```

