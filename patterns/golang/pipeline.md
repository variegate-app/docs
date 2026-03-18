# Pipeline

## [<<< ---](../index.md)

Паттерн **Pipeline** описывает последовательность стадий обработки данных, где каждая стадия:

- принимает значения из входного канала;
- выполняет свою часть работы;
- отправляет результат в выходной канал.

Каждая стадия может выполняться в отдельной горутине, что позволяет распараллелить обработку.

```go
// Stage — тип функции стадии конвейера.
type Stage func(<-chan int) <-chan int

// MultiplyStage возвращает стадию, умножающую все входные значения на factor.
func MultiplyStage(done <-chan struct{}, factor int) Stage {
	return func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for v := range in {
				select {
				case <-done:
					return
				case out <- v * factor:
				}
			}
		}()
		return out
	}
}

// AddStage возвращает стадию, добавляющую к каждому значению inc.
func AddStage(done <-chan struct{}, inc int) Stage {
	return func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for v := range in {
				select {
				case <-done:
					return
				case out <- v + inc:
				}
			}
		}()
		return out
	}
}

// BuildPipeline последовательно применяет стадии к начальному каналу.
func BuildPipeline(in <-chan int, stages ...Stage) <-chan int {
	out := in
	for _, s := range stages {
		out = s(out)
	}
	return out
}
```

Такой конвейер легко расширять и переиспользовать, добавляя или переставляя стадии.

