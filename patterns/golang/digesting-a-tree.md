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

