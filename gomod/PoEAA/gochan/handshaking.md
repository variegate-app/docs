# Handshaking

## [<<< ---](../gochan.md)

Паттерн **Handshaking** описывает явное согласование (рукопожатие) между двумя сторонами при помощи каналов: отправитель не только посылает данные, но и ждёт подтверждения приёма.

Это полезно, когда нужно быть уверенным, что сообщение не только ушло в канал, но и было обработано.

```go
type Message struct {
	Payload string
	Ack     chan struct{} // канал для подтверждения
}

func sender(out chan<- Message, payload string) {
	ack := make(chan struct{})
	out <- Message{Payload: payload, Ack: ack}

	// ждём подтверждения обработки
	<-ack
}

func receiver(in <-chan Message) {
	for msg := range in {
		fmt.Println("получено:", msg.Payload)
		// ... обработка ...

		// подтверждаем получателю, что событие обработано
		close(msg.Ack)
	}
}
```

Такой «handshake» может использоваться, например, для надёжной доставки команд в управляющий цикл или FSM.

