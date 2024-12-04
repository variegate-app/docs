# or-channel

## [<<< ---](../gochan.md)

This pattern aims to combine multiple `done` channels into one `agg_done`; it means that if one of a `done` channel is signaled, the whole `agg_done` channel is also closed. Yet, we do not know number of `done` channels during runtime in advanced.

`or-channel` pattern can do so by using `goroutine` & `recursion` .

```go
// return agg_done channel
var or func(channels ... <-chan interface{}) <- chan interface{}

or = func(channels ...<-chan interface{}) <-chan interface{} {
    // base cases
    switch len(channels) {
        case 0: return nil
        case 1: return channels[0]
    }

    orDone := make(chan interface{})

    go func() {
        defer close(orDone)

        switch len(channels) {
            case 2:
                select {
                    case <- channels[0]:
                    case <- channels[1]:
                }
            default:
                select {
                    case <- channels[0]:
                    case <- channels[1]:
                    case <- channels[2]:
                    case <- or(append(channels[3:], orDone)...): // * line
                }

        }

    }
    return orDone
}

```

line * makes the upper & lower recursive function depends on each other like a tree. The upper injects its own `orDone` channel into the lower. Then the lower also return its own `orDone` to the upper.

If any `orDone` channel closes, the upper & lower both are notified.