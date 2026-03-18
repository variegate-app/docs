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