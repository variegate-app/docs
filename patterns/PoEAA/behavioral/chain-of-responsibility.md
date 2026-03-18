# Цепочка обязанностей (Chain of Responsibility)

## [<<< ---](../../index.md)

## Назначение

Цепочка обязанностей позволяет передавать запрос по цепочке обработчиков до тех пор, пока один из них не сможет обработать запрос.

Обработчики не знают конкретного получателя «напрямую» — они либо обрабатывают, либо переадресуют дальше.

## Когда использовать

- когда набор возможных обработчиков неизвестен заранее или может меняться;
- когда нужно разделить логику обработки по уровням/условиям;
- когда требуется гибкая последовательность действий.

## Пример реализации на Go

```go
package main

import "fmt"

type Handler interface {
	SetNext(Handler)
	Handle(req int)
}

type base struct{ next Handler }

func (b *base) SetNext(h Handler) { b.next = h }

type EvenHandler struct{ base }

func (h *EvenHandler) Handle(req int) {
	if req%2 == 0 {
		fmt.Println("even:", req)
		return
	}
	if h.next != nil {
		h.next.Handle(req)
	}
}

type PositiveHandler struct{ base }

func (h *PositiveHandler) Handle(req int) {
	if req > 0 {
		fmt.Println("positive:", req)
		return
	}
	if h.next != nil {
		h.next.Handle(req)
	}
}

func main() {
	h1 := &EvenHandler{}
	h2 := &PositiveHandler{}
	h1.SetNext(h2)

	for _, v := range []int{2, -3, 5, 0} {
		h1.Handle(v)
	}
}
```

