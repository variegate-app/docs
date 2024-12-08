# bridge channel

## [<<< ---](../gochan.md)

Reading values from channel of channels (`<-chan <-chan interface{}`) can be cumbersome. Hence, this pattern aims to merge all values into 1 channel, so that the consumer jobs is much easier.

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