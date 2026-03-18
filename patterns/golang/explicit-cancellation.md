# Explicit cancellation

## [<<< ---](../index.md)

Паттерн **Explicit cancellation** подчёркивает явную, а не неявную отмену работы горутин: вызывающий код в явном виде посылает сигнал отмены, а все участники конвейера обязаны его обрабатывать.

В Go это реализуется через:

- отдельный канал `done`;
- или `context.Context`.

```go
func worker(done <-chan struct{}, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-done:
			// получен явный сигнал отмены
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			results <- j * 2
		}
	}
}

func main() {
	done := make(chan struct{})
	jobs := make(chan int)
	results := make(chan int)

	go worker(done, jobs, results)

	// отправляем несколько задач
	for i := 0; i < 3; i++ {
		jobs <- i
	}

	// явно отменяем дальнейшую работу
	close(done)
	close(jobs)
}
```

Явная отмена позволяет избежать утечек горутин и «подвисания» каналов.

