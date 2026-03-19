# bridge channel

## [<<< ---](../index.md)

Чтение значений из «канала каналов» (`<-chan <-chan interface{}`) может быть неудобным. Этот паттерн позволяет объединить все значения в один канал, чтобы потребителю было проще с ними работать.

```go
bridge := func(
    done <- chan interface{},
    chanStream <- <- interface{},
) <- chan interface{} {
    valStream := make(chan interface{})

    go func() {
        defer close(valStream)

        for {
            var stream <- chan interface{}
            select {
            case maybeStream, ok := <-chanStream
                if ok == false {
                    return
                }
                stream = maybeStream
            case <- done:
                return
            }

            for val := range orDone(done, stream){
                select{
                case valStream <- val:
                case <- done:
                }
            }
        }
    }()
    return valStream
}
```

## Типовые ошибки / антипаттерны
- Забывать, что “канал каналов” может содержать закрытые/пустые каналы: нужно обрабатывать `ok == false`.
- Отсутствие отмены: если `done` не слушается во вложенном цикле, bridge “зависнет” на чтении.
- Не закрывать выходной канал: потребитель будет ждать `out` навсегда.

## Практический чеклист
- Внутренние goroutine закрывают `valStream` в единой точке.
- Для каждого входного stream используется корректный цикл чтения до закрытия.
- Есть единая отмена (`done`) на всех уровнях чтения (outer + inner).

## Как адаптировать под кейс
- Bridge особенно полезен, когда источники потоков появляются динамически (fan-in каналов).
- Если каналов очень много — проверьте стоимость горутин и переключений.

## Связанные паттерны
- `./or-channel.md`
- `./for-select-done.md`
- `./fan-in.md`