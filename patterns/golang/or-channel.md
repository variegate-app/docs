# or-channel

## [<<< ---](../index.md)

Этот паттерн предназначен для объединения нескольких каналов `done` в один общий `agg_done`. Это означает, что как только любой из каналов `done` получает сигнал (закрывается), общий канал `agg_done` тоже закрывается. При этом количество каналов `done` заранее во время выполнения может быть неизвестно.

Паттерн `or-channel` реализуется с помощью горутин и рекурсии.

```go
// возвращает агрегирующий канал agg_done
var or func(channels ...<-chan interface{}) <-chan interface{}

or = func(channels ...<-chan interface{}) <-chan interface{} {
    // базовые случаи
    switch len(channels) {
    case 0:
        return nil
    case 1:
        return channels[0]
    }

    orDone := make(chan interface{})

    go func() {
        defer close(orDone)

        switch len(channels) {
        case 2:
            select {
            case <-channels[0]:
            case <-channels[1]:
            }
        default:
            select {
            case <-channels[0]:
            case <-channels[1]:
            case <-channels[2]:
            case <-or(append(channels[3:], orDone)...): // * строка
            }

        }

    }()
    return orDone
}

```

Строка, отмеченная `*`, делает так, что верхний и нижний рекурсивные вызовы зависят друг от друга, образуя дерево: верхний передаёт свой канал `orDone` во внутренний вызов, а тот, в свою очередь, возвращает свой `orDone` наружу.

Если какой‑либо канал `orDone` закрывается, об этом уведомляются и верхние, и нижние уровни рекурсии.