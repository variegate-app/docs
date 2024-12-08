# Pooling

## [<<< ---](../gochan.md)

The main idea behind **Pooling** pattern is to have:

- a channel that provides a signaling semantics
    - unbuffered channel is used to have a guarantee a goroutine has received a signal
- multiple goroutines that pool that channel for work
- a goroutine that sends work via channel

### Example

In this example you are a `manager`, and you hire a bunch of new `employees`.

`Employees` don't know immediately what do to, and they wait for `manager` to give them some work. The are looking at the channel `ch` to see if there is some work to do.

Once `manager` finds some work for the `employees`, it notifies them by sending a signal (`paper`) via communication channel `ch`.

First available `employee` that sees a signal from the channel `ch`, takes and completes the work.

After that `employee` completes the work, he is once again available to do more work, and he starts waiting for a new signal on channel `ch`.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // make channel of type string which provides signaling semantics
    // unbuffered channel provides a guarantee that the
    // signal being sent is received
    ch := make(chan string)

    // number of goroutines to create, numCPU() is a good starting point
    //g := runtime.NumCPU()
    g := 3

    for e := 0; e < g; e++ {
        // a new goroutine is created for each employee
        go func(emp int) {
            // employee waits for the signal that there is some work to do
                        // all goroutines are blocked on the same channel `ch` recieve
            for p := range ch {
                fmt.Printf("employee %d : received signal : %s\n", emp, p)
            }

            // when all work is sent, manager notifies all employees by closing the channel
            // once the channel is closed, employee breaks out of the for-range loop
            fmt.Printf("employee %d : revieved shutdown signal\n", emp)
        }(e)
    }

    // amount of work to be done
    const work = 10

    for w := 0; w < work; w++ {
        // when work is ready, we send signal from the manager to the employee
        // sender (manager) has a guarantee that the worker (employee) has received the signal
        // manager doesn't care about which employee received a signal,
        // since all employees are capable of doing the work
        ch <- "paper"

        fmt.Println("manager : sent signal :", w)
    }

    // when all work is sent the manages notifies all employees by closing the channel
    // unbuffered channel provides a guarantee that all work has been sent
    close(ch)
    fmt.Println("manager : sent shutdown signal")

    time.Sleep(time.Second)

}

```

### Result (1st execution)

```
go run main.go

employee 2 : recieved signal : paper
manager : sent signal : 0
manager : sent signal : 1
manager : sent signal : 2
manager : sent signal : 3
employee 1 : recieved signal : paper
employee 1 : recieved signal : paper
employee 2 : recieved signal : paper
manager : sent signal : 4
manager : sent signal : 5
manager : sent signal : 6
employee 1 : recieved signal : paper
employee 1 : recieved signal : paper
employee 0 : recieved signal : paper
employee 2 : recieved signal : paper
manager : sent signal : 7
manager : sent signal : 8
manager : sent signal : 9
manager : sent shutdown signal
employee 0 : recieved signal : paper
employee 0 : revieved shutdown signal
employee 2 : revieved shutdown signal
employee 1 : recieved signal : paper
employee 1 : revieved shutdown signal
```