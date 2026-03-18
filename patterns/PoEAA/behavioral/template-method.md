# Шаблонный метод (Template Method)

## [<<< ---](../../index.md)

## Назначение

Template Method задаёт скелет алгоритма в базовом классе, позволяя подклассам переопределять отдельные шаги.

Это сохраняет общий порядок выполнения, но делает конкретные шаги расширяемыми.

## Когда использовать

- когда порядок шагов алгоритма почти всегда одинаков, но детали могут различаться;
- когда хотите вынести общий код и минимизировать дублирование между реализациями;
- когда требуется контролируемая специализация.

## Пример реализации на Go

```go
package main

import "fmt"

// TemplateMethod задает скелет алгоритма.
type Template interface {
	Step1()
	Step2()
	Step3()
}

func RunTemplate(t Template) {
	t.Step1()
	t.Step2()
	t.Step3()
}

type A struct{}

func (A) Step1() { fmt.Println("A: step1") }
func (A) Step2() { fmt.Println("A: step2") }
func (A) Step3() { fmt.Println("A: step3") }

type B struct{}

func (B) Step1() { fmt.Println("B: step1") }
func (B) Step2() { fmt.Println("B: step2") }
func (B) Step3() { fmt.Println("B: step3") }

func main() {
	RunTemplate(A{})
	RunTemplate(B{})
}
```

