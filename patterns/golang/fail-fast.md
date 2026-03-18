# Fail-Fast

## [<<< ---](../index.md)

Паттерн **Fail-Fast** подразумевает немедленное прекращение работы, как только обнаружена критическая ошибка, вместо продолжения вычислений в заведомо некорректном состоянии.

Для конкурентного кода это часто выглядит как:

- первый воркер, столкнувшийся с ошибкой, отправляет её в общий канал ошибок;
- управляющая горутина отменяет контекст и останавливает остальные воркеры.

```go
func processAll(ctx context.Context, inputs []string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errs := make(chan error, 1)
	var wg sync.WaitGroup

	for _, in := range inputs {
		wg.Add(1)
		go func(input string) {
			defer wg.Done()
			if err := doOne(ctx, input); err != nil {
				// отправляем только первую ошибку
				select {
				case errs <- err:
					cancel() // даём сигнал всем остальным остановиться
				default:
				}
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	return <-errs
}
```

Такой подход уменьшает нагрузку на систему и ускоряет обнаружение критических сбоёв.

