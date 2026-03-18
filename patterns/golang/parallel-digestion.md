# Parallel digestion

## [<<< ---](../index.md)

Паттерн **Parallel digestion** дополняет «Digesting a tree», добавляя параллельную обработку узлов дерева (например, файлов) с помощью пула горутин.

```go
func digestFiles(done <-chan struct{}, paths <-chan string, workers int) <-chan FileDigest {
	out := make(chan FileDigest)
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for path := range paths {
				data, err := os.ReadFile(path)
				select {
				case <-done:
					return
				case out <- FileDigest{
					Path: path,
					Sum:  md5.Sum(data),
					Err:  err,
				}:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```

Сначала мы строим поток путей (`walkTree`), затем параллельно «перевариваем» (digest) каждый файл и собираем результаты через один выходной канал.

