# Select Statement with Timeout

## [<<< ---](../gochan.md)

Timeout in select can be achieved by using **After()** function of time package. Below  is the signature of **After()** function.

```go
func After(d Duration) <-chan Time
```

The **After** function waits for d duration  to finish and then it returns the current time on a channel

[https://golang.org/pkg/time/#Time.After](https://golang.org/pkg/time/#Time.After)

# **Code**

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go goOne(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

func goOne(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "From goOne goroutine"
}
```

**Output**

```go
Timeout
```

In the above select statement we are waiting for for the receive operation to complete on **ch1**. In other case statement we have **time.After** with the duration of 1 second. So essentially this select statement will wait for at least 1 second for receive operation to complete on **ch1** and after that the **time.After** case statement will be executed.  We have put a timeout of more than 1 seconds in the **goOne** function and hence we  see the **time.After** statement getting executed and

```go
Timeout
```

getting printed as output.

So time.After() is a channel operation that gets unblocked after some time.

# **Timeout with infinite for loop outside select statement**

We can have infinite for loop outside a select statement. This will cause the select statement to execute indefinite number of times.So when using for statement with infinite loop outside the select statement, we need to have a way yo break out of the for loop. One of the use case of having infinite for look outside select statement could be that you are waiting for multiple operations to receive on a particular channel for a certain time.

See below example

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	news := make(chan string)
	go newsFeed(news)

	printAllNews(news)
}

func printAllNews(news chan string) {
	for {
		select {
		case n := <-news:
			fmt.Println(n)
		case <-time.After(time.Second * 1):
			fmt.Println("Timeout: News feed finished")
			return
		}
	}
}

func newsFeed(ch chan string) {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Millisecond * 400)
		ch <- fmt.Sprintf("News: %d", i+1)
	}
}
```

**Output**

```go
News: 1
News: 2
Timeout: News feed finished
```

In the above program, we have created a channel named **news**  which will hold data of string type. Then we pass this channel to the **newsfeed** function which is pushing the news feed to this channel . In the select statement, we are receiving the news feed from the **news** channel. This select statement is inside an infinite for loop  so the select statement will be executed multiple times until we  exit out of for loop . We also have **time.After** with a duration for 1 second as one of the case statements. So this set up will receive all the news from the **news** channel for  1  second and then exit.