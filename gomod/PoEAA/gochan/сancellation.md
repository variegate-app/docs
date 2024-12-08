# Cancellation

## [<<< ---](../gochan.md)

The main idea behind the **Cancellation Pattern** is to have a limited amount of time to perform work. If limit is reached, the work is ignored.

We have:

- a context with specified timeout
- a buffered channel that provides signaling semantic
- a worker goroutine that does the work
- a manager goroutine that waits on (which comes first):
    - worker goroutine signal (that the work is completed)
    - context timeout signal

### Example

In **Cancellation Pattern** we have a limited amount of time to perform some work.

Imagine we are in the ice cream making business and we have:

- a `manager` (`main` goroutine) that get and holds a scoop of the ice cream in one hand
    - he holds out his other hand (`communication channel`) and waits for an `employee` (`worker` goroutine) to pass him the ice cream cone so he can sell the ice cream
- an `employee` (`worker` goroutine) that needs some time to get the ice cream cone so he can pass it to the mangers' `hand`
    - if `employee` takes too much time to get the ice cream cone, the ice cream that `manager` holds will melt, so he won't need the ice cream cone anymore and the manager won't hold his hand anymore
- an `employee` is now stuck with the ice cream cone in his hand and can't perform any other work (`goroutine leak`)
    - to fix this `employee` and the `manager` decided to use new `communication` channel (e.g. desk `buffered channel`) so that `employee` can complete his work, regardless of the managers' hand

### Use Case

Good use case for this pattern is any request to a remote service, e.g. database request, API request or whatever request that can block. Since we don't want our request to block forever, we use timeout to cancel it.

Feel free to try the example on [Go Playground](https://play.golang.com/p/RCy0Iajt0tl)

```go
package main

import (
    "context"
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // a duration that sets the max time to perform the operation

    // e.g 150 ms to get the ice cream cone
    duration := 1 * time.Millisecond

    // context.Background() returns a non-nil, empty Context.
    // It is never canceled, has no values, and has no deadline.
    // It is typically used by the main function, initialization, and tests,
    // and as the top-level Context for incoming requests.
    emptyCtx := context.Background()

    // Create new context from emptyCtx + add timeout of 150 ms
    // ctx is new context with timeout
    // cancel is a function that releases resources associated with context

    // e.g. ticker that the manager uses to check if he gets the ice cream cone fast enough
    ctx, cancel := context.WithTimeout(emptyCtx, duration)

    // Canceling this context releases resources associated with it,
    // so code should call cancel as soon as the operations running in
    // this Context complete

    // e.g. command manager uses to cancel the context (unit of work - getting ice cream cone)
    defer cancel()

    // IMPORTANT:
    // Make buffered channel of size 1, and type string which provides signaling semantics.
    // Buffered channel ensures that the worker goroutine can perform the send operation
    // and complete even if there is no-one on the receive side.
    // e.g.
    // - if worker goroutine does NOT finish in 150ms
    // -- main goroutine will continue
    // -- this will cause worker goroutine leak
    //    since there is no-one goroutine to receive the sent signal (so it blocks and waits)

    // e.g. used to prevent employee from being blocked if he doesn't complete the work in 150ms
    ch := make(chan string, 1)

    // create worker goroutine
    go func() {
        // Simulate the idea of unknown latency (do not use in production).
        // Don't forget that context timeout is 150 ms, but this can take up to 200 ms.

        // e.g. employee reaches out for the ice cream cone from the box
        time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

        // send signal when work is done
        // e.g. employee passes ice creeam cone to the managers' hand
        ch <- "paper"
    }()

    // select-case allow us to perform multiple channel operations
    // at the same time, on the same goroutine

    // e.g. manager waits for ice cream cone, or for 150 ms timer to time out
    select {

    // best case scenario:
    // receive a result from worker goroutine in under the 150 ms

    // e.g. employee finds and passes the ice cream cone to the manager
    case d := <-ch:
        fmt.Println("work complete", d)

    // ctx.Done() call starts the 150ms duration clock ticking.
    // If 150 ms passes before the worker goroutine finishes, this println will be executed

    // e.g. manager doesn't wait for employee to get the ice cream cone anymore
    case <-ctx.Done():
        fmt.Println("work cancelled")
    }
}

```

### Result (1st execution)

```
go run main.go

work complete paper

```

### Result (2st execution)

```
go run main.go

work cancelled
```