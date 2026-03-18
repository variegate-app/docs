# Publish & Subscribe

## [<<< ---](../gochan.md)

Publish–Subscribe — это паттерн обмена сообщениями, позволяющий передавать сообщения между компонентами, не заставляя их знать что‑либо о личности друг друга.

Он похож на поведенческий паттерн Observer. Ключевая идея и Observer, и Publish–Subscribe — разделить отправителя событий (`Event Messages`) и всех, кто хочет о них узнавать (наблюдатели или подписчики). То есть нам не нужно жёстко прописывать, каким конкретным получателям надо отправлять сообщения.

Для этого используется посредник — «message broker» или «event bus», который принимает опубликованные сообщения и маршрутизирует их подписчикам.

Есть три сущности: **messages**, **topics**, **users**.

```go
type Message struct {
    // содержимое сообщения
}

type Subscription struct {
    ch chan<- Message

    Inbox chan Message
}

func (s *Subscription) Publish(msg Message) error {
    if _, ok := <-s.ch; !ok {
        return errors.New("Topic has been closed")
    }

    s.ch <- msg

    return nil
}

type Topic struct {
    Subscribers    []Session
    MessageHistory []Message
}

func (t *Topic) Subscribe(uid uint64) (Subscription, error) {
    // Найти (или создать) сессию для пользователя.

    // Добавить сессию в Topic и в MessageHistory.

    // Создать и вернуть Subscription.
}

func (t *Topic) Unsubscribe(Subscription) error {
    // Логика отписки.
}

func (t *Topic) Delete() error {
    // Логика удаления топика.
}

type User struct {
    ID uint64
    Name string
}

type Session struct {
    User User
    Timestamp time.Time
}
```

# **Improvements**

Events can be published in a parallel fashion by utilizing stackless goroutines.

Performance can be improved by dealing with straggler subscribers by using a buffered inbox and you stop sending events once the inbox is full.