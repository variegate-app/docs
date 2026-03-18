# Интерпретатор (Interpreter)

## [<<< ---](../../index.md)

## Назначение

Интерпретатор описывает грамматику (язык/выражения) как структуру объектов и реализует интерпретацию этой структуры.

Идея: преобразовать вход (AST) в набор объектов, а затем «исполнять» его рекурсивно.

## Когда использовать Interpreter

- когда нужно реализовать небольшой DSL (domain-specific language) или набор правил;
- когда правила меняются часто и их удобно описывать структурой;
- когда важна читаемость и расширяемость интерпретации.

## Пример реализации на Go

```go
package main

import "fmt"

// Expression — узел AST, который умеет вычисляться.
type Expression interface {
	Eval(vars map[string]int) int
}

// Number — числовая константа.
type Number struct{ v int }

func (n Number) Eval(_ map[string]int) int { return n.v }

// Variable — переменная.
type Variable struct{ name string }

func (v Variable) Eval(vars map[string]int) int { return vars[v.name] }

// Add — бинарная операция: a + b.
type Add struct {
	left, right Expression
}

func (a Add) Eval(vars map[string]int) int {
	return a.left.Eval(vars) + a.right.Eval(vars)
}

func main() {
	// (x + 10) + (y + 5)
	expr := Add{
		left: Add{left: Variable{"x"}, right: Number{10}},
		right: Add{left: Variable{"y"}, right: Number{5}},
	}

	res := expr.Eval(map[string]int{"x": 2, "y": 20})
	fmt.Println("result:", res)
}
```

