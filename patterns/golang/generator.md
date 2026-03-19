# Генератор

## [<<< ---](../index.md)

## **Реализация**

```go
func Count(start int, end int) chan int {
    ch := make(chan int)

    go func(ch chan int) {
        for i := start; i <= end ; i++ {
            // Блокируется на операции отправки
            ch <- i
        }

        close(ch)
    }(ch)

    return ch
}
```

## **Использование**

```go
fmt.Println("No bottles of beer on the wall")

for i := range Count(1, 99) {
    fmt.Println("Pass it around, put one up,", i, "bottles of beer on the wall")
    // Pass it around, put one up, 1 bottles of beer on the wall
    // Pass it around, put one up, 2 bottles of beer on the wall
    // ...
    // Pass it around, put one up, 99 bottles of beer on the wall
}

fmt.Println(100, "bottles of beer on the wall")
```

## Типовые ошибки / антипаттерны
- Генератор не уважает отмену/закрытие: после остановки consumers goroutine продолжает работать.
- Выходной канал не закрывается: pipeline зависает в ожидании значений.
- Слишком большой буфер генератора: теряется backpressure и растёт память.

## Практический чеклист
- Генератор закрывает выходной канал после завершения диапазона.
- В длительных генерациях есть проверка `done/ctx.Done()`.
- Генератор не делает лишнюю работу, если потребителю уже не нужны значения.

## Как адаптировать под кейс
- Генератор удобно использовать как источник для `pipeline`, `filter`, `map-filter`, `take-first-n`.

## Связанные паттерны
- `./pipeline.md`
- `./filter.md`
- `./take-first-n.md`