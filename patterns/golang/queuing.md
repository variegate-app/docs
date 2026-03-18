# Queuing

## [<<< ---](../index.md)

Паттерн **Queuing** использует буферизованный канал как простейшую очередь заданий между производителями и потребителями.

Буфер канала сглаживает всплески нагрузки: производители могут непродолжительно работать быстрее потребителей, пока очередь не заполнится.

```go
// Producer публикует задачи в очередь queue.
func Producer(queue chan<- string, items []string) {
	for _, it := range items {
		queue <- it
	}
}

// Consumer читает задачи из очереди и обрабатывает их.
func Consumer(done <-chan struct{}, queue <-chan string) {
	for {
		select {
		case <-done:
			return
		case item, ok := <-queue:
			if !ok {
				return
			}
			// обработка item
			fmt.Println("processed:", item)
		}
	}
}

func main() {
	done := make(chan struct{})
	defer close(done)

	// очередь с буфером на 100 элементов
	queue := make(chan string, 100)

	go Consumer(done, queue)
	Producer(queue, []string{"a", "b", "c"})
	close(queue)
}
```

При необходимости поведение можно усложнить: добавить несколько потребителей, приоритеты, ретраи и т.д.

