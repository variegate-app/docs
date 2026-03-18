# Посетитель (Visitor)

## [<<< ---](../../index.md)

## Назначение

Visitor отделяет алгоритмы от структуры объектов.

Он позволяет добавлять новые операции над объектами, не меняя классы самой структуры (часто — за счёт двойной диспетчеризации).

## Когда использовать Visitor

- когда структура объектов относительно стабильна, но операций над ними появляется много;
- когда нужно поддерживать разные операции для разных типов элементов структуры;
- когда логика операций сильно зависит от типа элемента.

## Пример реализации на Go

```go
package main

import "fmt"

// Element — часть структуры, которая принимает посетителя.
type Element interface {
	Accept(v Visitor)
}

type Visitor interface {
	VisitBook(*Book)
	VisitMovie(*Movie)
}

type Book struct{ Title string }

func (b *Book) Accept(v Visitor) { v.VisitBook(b) }

type Movie struct{ Name string }

func (m *Movie) Accept(v Visitor) { v.VisitMovie(m) }

type PrintVisitor struct{}

func (p PrintVisitor) VisitBook(b *Book)  { fmt.Println("book:", b.Title) }
func (p PrintVisitor) VisitMovie(m *Movie) { fmt.Println("movie:", m.Name) }

func main() {
	var elems []Element = []Element{
		&Book{Title: "Go in Action"},
		&Movie{Name: "Gophers"},
	}

	visitor := PrintVisitor{}
	for _, e := range elems {
		e.Accept(visitor)
	}
}
```

