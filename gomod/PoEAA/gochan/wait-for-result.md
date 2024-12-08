# Wait For Result

## [<<< ---](../gochan.md)

The main idea behind **Wait For Result** pattern is to have:

- a channel that provides a signaling semantics
- a goroutine that does some work
- a goroutine that waits for that work to be done

### Example

In this example we have an `employee` (`a` goroutine) that has some work to do. We also have a `manager` (`main` goroutine) that waits on that work to be done. Once work is done, `employee` notifies `manager` by sending a signal (`paper`) via communication channel `ch`.

```go
package main

func main() {
    // make channel of type string which provides signaling semantics
    ch := make(chan string)

    // goroutine 'a' that does some work => employee
    go func() {
        // simulate the idea of unknown latency (do not use in production)
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
                // when work is done send signal
        ch <- "paper"

        // we don't know the order of print statement
        fmt.Println("employee: sent signal")
    }()

    // goroutine 'main' => manager
    // goroutines 'main' and 'a' are executed in parallel

    // wait for and receive signal from 'goroutine a'
        // blocking operation
    p := <-ch

    // we don't know which print statement is going to be executed first
    fmt.Println("manager: received signal :", p)

    // ensure enough time to get the result (do not use in production)
    time.Sleep(time.Second)
}

```

### Result (1st execution)

```go
go run main.go
employee: sent signal
manager: received signal : paper
```

### Result (2st execution)

```
go run main.go
manager: received signal : paper
employee: sent signal

```

## Most common misconceptions:

1. Note that in the first and the second execution of the previous code, the order of `fmt.Println()` statement is not guaranteed. Goroutines are executed in parallel.