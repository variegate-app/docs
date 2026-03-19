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

## Типовые ошибки / антипаттерны
- Ставить `workers` слишком большим без оценки: вы убьёте диск/CPU и получите деградацию.
- Не учитывать отмену: workers продолжают читать/вычислять после stop.
- Не закрывать out: consumer будет ждать закрытия бесконечно.

## Практический чеклист
- Количество workers ограничено и подбирается под ресурсы.
- Выходной канал `out` закрывается после `wg.Wait()`.
- В workers есть select-ветка на `<-done`.
- Ошибки (если они есть) либо передаются в `FileDigest.Err`, либо агрегируются upstream.

## Как адаптировать под кейс
- Для тяжелых вычислений используйте несколько стадий pipeline: чтение отдельно от хеширования/парсинга.
- Если важно ограничить IO concurrency — подберите отдельный limit на чтение.

## Связанные паттерны
- `./digesting-a-tree.md`
- `./bounded-parallelism.md`
- `./worker-pool.md`

