# Filter

## [<<< ---](../gochan.md)

Паттерн **Filter** используется для отсеивания ненужных значений из потока данных, проходящего через канал.  
Он часто является одним из шагов в большом конвейере обработки.

```go
// Filter пропускает только те значения, которые удовлетворяют предикату pred.
func Filter[T any](done <-chan struct{}, in <-chan T, pred func(T) bool) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for v := range in {
			if !pred(v) {
				continue
			}

			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()

	return out
}
```

Пример: фильтрация только положительных чисел из потока `ints`:

```go
positives := Filter(done, ints, func(v int) bool { return v > 0 })
```

