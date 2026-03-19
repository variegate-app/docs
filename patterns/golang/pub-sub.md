# Publish & Subscribe

## [<<< ---](../index.md)

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

## Типовые ошибки / антипаттерны
- Доставлять события синхронно всем подписчикам без буферов: медленный подписчик “прибивает” весь delivery.
- Отсутствие стратегии `unsubscribe`/cleanup: подписчики копятся, утечки памяти/каналов.
- Не определять поведение при переполнении inbox (drop vs block vs retry).

## Практический чеклист
- Сущности разделены: `Topic`/`Broker` управляет списком подписчиков и жизненным циклом.
- Доставка событий учитывает backpressure (буфер inbox, drop политика, ограничение скорости).
- Подписчики получают события через канал, и их горутины корректно завершаются при shutdown.
- Есть идемпотентность/поведение на дубликаты сообщений (если delivery может повторяться).

## Как адаптировать под кейс
- Для “важно не потерять” выбирайте очередь/retention + подтверждения; для “важно не блокировать” применяйте drop при переполнении.
- Настройте размер buffered inbox под измеренную латентность подписчиков.

## Связанные паттерны
- `./context.md`
- `./graceful-shutdown.md`
- `../PoEAA/behavioral/observer.md`