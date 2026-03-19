# Map & Filter

## [<<< ---](../index.md)

Паттерны **Map** и **Filter** переносят привычные функциональные операции на коллекции в мир потоков и каналов:

- **Map** преобразует каждое входное значение по заданной функции и отдаёт результат в выходной канал;
- **Filter** пропускает или отбрасывает значения в зависимости от предиката.

Эти операции хорошо комбинируются внутри конвейера (pipeline).

```go
// Map применяет функцию mapper ко всем значениям из in и отправляет результат в новый канал.
func Map[T any, R any](done <-chan struct{}, in <-chan T, mapper func(T) R) <-chan R {
	out := make(chan R)

	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
				return
			case out <- mapper(v):
			}
		}
	}()

	return out
}

// Filter пропускает только те значения, для которых pred(v) == true.
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

Пример использования:

```go
// src -> фильтруем только чётные -> возводим в квадрат
ints := generator(done, 1, 2, 3, 4, 5, 6)
evens := Filter(done, ints, func(v int) bool { return v%2 == 0 })
squares := Map(done, evens, func(v int) int { return v * v })
```

## Типовые ошибки / антипаттерны
- Не закрывать выходные каналы из map/filter: конвейер зависает.
- Делать long-running/блокирующую работу прямо внутри mapper/pred без отдельной стадии.
- Игнорировать `done`: отмена не остановит цепочку.

## Практический чеклист
- Для `Map` и `Filter` выходные каналы закрываются после завершения входа/отмены.
- `done` проверяется в `select` перед записью в `out`.
- `mapper/pred` идемпотентны или имеют понятную политику повторов (если upstream будет перезапускать pipeline).

## Как адаптировать под кейс
- Комбинируйте `Map`/`Filter` как функциональные стадии: это повышает читаемость конвейера.

## Связанные паттерны
- `./pipeline.md`
- `./filter.md`
- `./generator.md`

