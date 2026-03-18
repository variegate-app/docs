# Фасад (Facade)

## [<<< ---](../../index.md)

## Назначение

Фасад предоставляет простой унифицированный интерфейс к более сложной системе классов, скрывая внутреннюю детализацию.

Он помогает уменьшить связанность клиентского кода с «мелкими» подсистемами.

## Когда использовать Facade

- когда подсистема слишком сложна и клиентам нужен «единый вход»;
- когда вы хотите сократить число точек взаимодействия с разными объектами;
- когда необходимо сохранить совместимость интерфейса, не распространяя изменения внутренней логики.

## Пример реализации на Go

```go
package main

import "fmt"

// Подсистема A
type EmailService struct{}

func (e EmailService) Send(to, subject string) { fmt.Println("email:", to, subject) }

// Подсистема B
type SMSService struct{}

func (s SMSService) Send(to, text string) { fmt.Println("sms:", to, text) }

// Фасад скрывает детали подсистем.
type Notifier struct {
	email EmailService
	sms   SMSService
}

func (n Notifier) NotifyUser(userID string) {
	// В реальном коде выбор подсистемы мог бы зависеть от настроек.
	n.email.Send(userID, "Welcome!")
	n.sms.Send(userID, "Your account is ready")
}

func main() {
	n := Notifier{email: EmailService{}, sms: SMSService{}}
	n.NotifyUser("user-1")
}
```

