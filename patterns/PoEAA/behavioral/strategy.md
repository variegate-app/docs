# Стратегия (Strategy)

## [<<< ---](../../index.md)

## Назначение

Strategy инкапсулирует алгоритм в отдельный объект и делает его заменяемым во время работы программы.

Это позволяет легко выбирать и менять поведение без изменения клиентского кода.

## Когда использовать

- когда есть несколько алгоритмов для одной задачи и выбор нужно делать динамически;
- когда алгоритмы похожи по интерфейсу, но отличаются деталями реализации;
- когда нужно отделить «что делаем» от «как делаем».

## Пример реализации на Go

```go
package main

import "fmt"

type SortStrategy interface {
	Sort([]int) []int
}

type Asc struct{}

func (a Asc) Sort(s []int) []int {
	out := append([]int(nil), s...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j] < out[i] {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}

type Desc struct{}

func (d Desc) Sort(s []int) []int {
	out := Asc{}.Sort(s)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}

type Context struct {
	strat SortStrategy
}

func (c Context) Sort(s []int) []int { return c.strat.Sort(s) }

func main() {
	s := []int{3, 1, 2}
	fmt.Println(Context{strat: Asc{}}.Sort(s))
	fmt.Println(Context{strat: Desc{}}.Sort(s))
}
```

