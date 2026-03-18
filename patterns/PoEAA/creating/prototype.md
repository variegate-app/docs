# Прототип (Prototype)

## [<<< ---](../../index.md)

## Назначение

Prototype клонирует объекты вместо того, чтобы создавать их заново через конструкторы.

Это особенно удобно, когда создание объекта дорого или когда конфигурация экземпляров разнообразна.

## Когда использовать

- когда нужно быстро создавать новые объекты «по шаблону»;
- когда создание через конструктор слишком сложное/дорогое;
- когда удобно копировать состояние и немного модифицировать клон.

## Пример реализации на Go

```go
package main

import "fmt"

// Prototype — объект, который умеет клонироваться.
type Prototype interface {
	Clone() Prototype
}

type User struct {
	ID   int
	Name string
}

func (u User) Clone() Prototype {
	// Для простых типов достаточно “поверхностного” копирования.
	c := u
	return &c
}

func main() {
	original := &User{ID: 1, Name: "Alice"}
	cloned := original.Clone().(*User)

	cloned.Name = "Alice (copy)"
	fmt.Println("original:", original.Name)
	fmt.Println("cloned:", cloned.Name)
}
```

