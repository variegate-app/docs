# Wait For Task

## [<<< ---](../gochan.md)

The main idea behind **Wait For Task** pattern is to have:

- a channel that provides a signaling semantics
- a goroutine that **waits for task** so it can do some work
- a goroutine that sends work to the previous goroutine

### Example

In this example we have an `employee` (`a` goroutine) that doesn't know immediately what to do. The `employee` waits for `manager` to give him some work to do.

Once `manager` finds some work for the `employee`, it notifies `employee` by sending a signal (`paper`) via communication channel `ch`.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // make channel of type string which provides signaling semantics
    // unbuffered channel provides a guarantee that the
    // signal being sent is received
    ch := make(chan string)

    // goroutine 'a' that waits for some work => employee
    go func() {
        // employee waits for signal that it has some work to do
        p := <-ch
        fmt.Println("employee : received signal : ", p)
    }()

    // simulate the idea of unknown latency (do not use in production)
    // e.g. manager is thinking what work to pass to the employee
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

    // when work is ready, send signal form manager to the employee
    // sender (employee) has a guarantee that the worker (employee)
    // has received a signal
    ch <- "paper"

    fmt.Println("manager : sent signal")

    time.Sleep(time.Second)
}
```

### Result (1st execution)

`go run main.go
manager : sent signal
employee : received signal :  paper`