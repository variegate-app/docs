# Адаптер (Adapter)

## [<<< ---](../../index.md)

## Назначение

Адаптер позволяет использовать несовместимые интерфейсы, оборачивая один объект так, чтобы он выглядел как другой.

Идея: «адаптируем» внешний или унаследованный API под интерфейс, нужный вашей системе.

## Когда использовать Adapter

- нужно подключить библиотеку/сервис, интерфейс которого не совпадает с вашими ожиданиями;
- вы хотите изолировать изменения внешнего API внутри одного класса-адаптера;
- важно сохранить существующую логику, не меняя её под новый интерфейс.

## Пример реализации на Go

```go
package main

import "fmt"

// Target — ожидаемый интерфейс в вашей системе.
type Target interface {
	Request() string
}

// Adaptee — компонент с несовместимым интерфейсом.
type Adaptee struct{}

func (a *Adaptee) DifferentRequest() string { return "hello from adaptee" }

// Adapter — оборачивает Adaptee и приводит его API к Target.
type Adapter struct {
	a *Adaptee
}

func (ad *Adapter) Request() string {
	return ad.a.DifferentRequest()
}

func main() {
	var t Target = &Adapter{a: &Adaptee{}}
	fmt.Println(t.Request())
}
```

