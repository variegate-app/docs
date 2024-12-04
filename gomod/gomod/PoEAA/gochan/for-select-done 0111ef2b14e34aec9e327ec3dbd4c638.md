# for-select-done

## [<<< ---](../gochan.md)

Goroutine is not garbage collected; hence, it is likely to be leaked.

```go
go func() {
// <operation that will block forever>
// => Go routine leaks
}()
// Do work

```

To avoid leaking, Goroutine should be cancelled whenever it is told to do. A parent Goroutine needs to send cancellation signal to its child via a *read-only*channel named `done` . By convention, it is set as the 1st parameter.

This pattern is also utilized a lot in other patterns.

```go
//child goroutine
doWork(<-done chan interface {}, other_params) <- terminated chan interface{} {
    terminated := make(chan interface{}) // to tell outer that it has finished
    defer close(terminated)

    for {
        select: {
            case: //do your work here
            case <- done:
                return
        }
        // do work here
    }

    return terminated
}

// parent goroutine
done := make(chan interface{})
terminated := doWork(done, other_args)

// do sth
// then tell child to stop
close (done)

// wait for child finish its work
<- terminated
```