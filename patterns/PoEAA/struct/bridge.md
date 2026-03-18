# Мост (Bridge)

## [<<< ---](../../index.md)

## Назначение

Мост разделяет абстракцию и реализацию так, чтобы они могли развиваться независимо.

Вместо иерархий «абстракция×реализация» вы получаете две иерархии: одну — для доменной абстракции, вторую — для реализации (платформы, среды, варианта поведения).

## Когда использовать Bridge

- когда у вас несколько измерений изменяемости (например, тип абстракции и платформа/вариант реализации);
- когда вы хотите уменьшить количество комбинаций в коде;
- когда требуется «контролируемая» замена реализации без переписывания абстракции.

## Пример реализации на Go

```go
package main

import "fmt"

// Implementor — реализация, которую “держит” абстракция.
type Implementor interface {
	DrawCircle(radius int)
}

type RedRenderer struct{}

func (r *RedRenderer) DrawCircle(radius int) {
	fmt.Printf("Draw RED circle r=%d\n", radius)
}

type BlueRenderer struct{}

func (b *BlueRenderer) DrawCircle(radius int) {
	fmt.Printf("Draw BLUE circle r=%d\n", radius)
}

// Abstraction — “верх” моста, которая использует Implementor.
type Shape struct {
	impl Implementor
}

func (s *Shape) Draw(radius int) {
	s.impl.DrawCircle(radius)
}

func main() {
	var shape Shape
	shape.impl = &RedRenderer{}
	shape.Draw(10)

	shape.impl = &BlueRenderer{}
	shape.Draw(7)
}
```

