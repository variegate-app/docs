# for-select

## [<<< ---](../gochan.md)

Один из базовых паттернов. Как правило используется для чтения данных из нескольких каналов

```go
var c1, c2 <-chan int

for { // Either loop infinitely or range over something
    select {
    case <-c1: // Do some work with channels
    case <-c2:
    default: // auto run if other cases are not ready
    }

    // do some work
}
```

The select statement looks like switch one, but its behavior is different. All `cases` are considered simultaneously & have **equal chance** to be selected. If none of the `cases` are ready to run, the entire `select` statement blocks.