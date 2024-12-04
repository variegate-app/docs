# tee channel

## [<<< ---](../gochan.md)

This pattern aims to split values coming from a channel into 2 others. So that we can dispatch them into two separate areas of our codebase.

```go
tee := func(
    done <- chan interface{},
    in <- chan interface{},
) (<- chan interface, <- chan interface) {
    out1 := make(chan interface{})
    out2 := make(chan interface{})

    go func() {
        defer close(out1)
        defer close(out2)

        //shadow outer variable
        var out1, out2 = out1, out2
        for val := range orDone(done, in) {
            for i := 0; i < 2; i ++ { //make sure 2 channels received same value
                select {
                case <- done:
                case out1<- val:
                    out1 = nil //stop this channel from being received
                case out2<-val:
                    out2 = nil
                }
            }
        }
    }()
    return out1, out2
}

```