# tee channel

## [<<< ---](../index.md)

Этот паттерн предназначен для разветвления значений из одного канала в два других, чтобы можно было направлять поток данных в две разные части кодовой базы.

```go
tee := func(
    done <-chan interface{},
    in <-chan interface{},
) (<-chan interface{}, <-chan interface{}) {
    out1 := make(chan interface{})
    out2 := make(chan interface{})

    go func() {
        defer close(out1)
        defer close(out2)

        // затеняем внешние переменные
        var out1, out2 = out1, out2
        for val := range orDone(done, in) {
            // гарантируем, что оба канала получат одно и то же значение
            for i := 0; i < 2; i++ {
                select {
                case <-done:
                    return
                case out1 <- val:
                    // перестаём писать в этот канал для текущего значения
                    out1 = nil
                case out2 <- val:
                    out2 = nil
                }
            }
        }
    }()
    return out1, out2
}

```