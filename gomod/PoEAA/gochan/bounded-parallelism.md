# Ограниченный параллелизм

## [<<< ---](../gochan.md)

```go
package bounded_parallelism

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

// walkFiles запускает горутину, которая обходит дерево каталогов root и
// отправляет путь каждого обычного файла в строковый канал. Результат обхода
// она отправляет в канал ошибок. Если канал done закрывается, walkFiles
// прекращает работу.
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() { // HL
		// Закрываем канал paths после завершения Walk.
		defer close(paths) // HL
		// Для этой отправки select не нужен, потому что errc буферизован.
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error { // HL
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path: // HL
			case <-done: // HL
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

// result — это результат чтения файла и вычисления его MD5-суммы.
type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// digester читает имена путей из канала paths и отправляет хеши соответствующих
// файлов в канал c, пока не будут закрыты либо paths, либо done.
func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths { // HLpaths
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

// MD5All читает все файлы в дереве каталогов с корнем root и возвращает карту
// из пути файла в MD5-сумму его содержимого. Если обход каталогов завершается
// с ошибкой или какая‑то операция чтения завершается с ошибкой, MD5All
// возвращает ошибку. В этом случае MD5All не дожидается завершения всех
// уже запущенных операций чтения.
func MD5All(root string) (map[string][md5.Size]byte, error) {
	// MD5All закрывает канал done при выходе; это может случиться до того,
	// как будут получены все значения из c и errc.
	done := make(chan struct{})
	defer close(done)

	paths, errc := walkFiles(done, root)

	// Запускаем фиксированное количество горутин для чтения и хеширования файлов.
	c := make(chan result) // HLc
	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c) // HLc
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c) // HLc
	}()
	// Конец конвейера. OMIT

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	// Проверяем, не завершился ли Walk ошибкой.
	if err := <-errc; err != nil { // HLerrc
		return nil, err
	}
	return m, nil
}

func main() {
	// Вычисляем MD5-сумму всех файлов под указанным каталогом
	// и печатаем результаты, отсортированные по имени пути.
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}

```