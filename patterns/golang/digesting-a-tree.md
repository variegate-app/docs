# Digesting a tree

## [<<< ---](../index.md)

Паттерн **Digesting a tree** относится к обходу и параллельной обработке древовидных структур (например, файловой системы).

Идея:

- обходим дерево (директории или структуру в памяти);
- для каждого листа/узла порождаем задание;
- результаты собираем через канал.

Ниже — упрощённый пример для файловой системы.

```go
type FileDigest struct {
	Path string
	Sum  [md5.Size]byte
	Err  error
}

func walkTree(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case <-done:
				return errors.New("walk cancelled")
			case paths <- path:
				return nil
			}
		})
	}()

	return paths, errc
}
```

Дальше поверх `paths` можно строить конвейер, который считает хеши файлов и агрегирует результаты.

## Типовые ошибки / антипаттерны
- Не закрыть каналы `paths`/`errc`: конвейер зависнет.
- Обход дерева не умеет отменяться по `done`: горутина продолжает работать после stop.
- Отдавать результат без backpressure: переполнение каналов вызывает рост памяти/задержек.

## Практический чеклист
- Для обхода есть отмена: `done` влияет на отправку в `paths`.
- Выходные каналы закрываются в единой точке (например, в goroutine-обёртке).
- Ошибки обрабатываются: `errc` возвращает ошибку корректно и не теряется.
- Downstream правильно закрывает pipeline после `done`/ошибки.

## Как адаптировать под кейс
- Для файловой системы используйте буферы/лимиты на количество одновременно открытых файлов (через parallel digestion).
- Для больших деревьев измеряйте стоимость Walk и размер параллелизма.

## Связанные паттерны
- `./parallel-digestion.md`
- `./pipeline.md`
- `./graceful-shutdown.md`

