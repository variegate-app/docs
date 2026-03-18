# Pooling

## [<<< ---](../gochan.md)

Основная идея паттерна **Pooling**:

- есть канал, предоставляющий семантику сигналов;
    - для гарантии приёма сигнала используется небуферизованный канал;
- есть несколько горутин, которые «слушают» этот канал в ожидании работы;
- есть горутина, которая отправляет работу через этот канал.

### Пример

В этом примере вы — `manager`, нанимающий пачку новых `employees`.

`Employees` поначалу не знают, что делать, и ждут, пока `manager` даст им работу. Они смотрят в канал `ch`, чтобы увидеть, появилась ли работа.

Как только `manager` находит задание для `employees`, он уведомляет их, отправляя сигнал (`paper`) через канал `ch`.

Первый свободный `employee`, который увидит сигнал в `ch`, забирает и выполняет работу.

После завершения работы `employee` снова становится доступен для новых задач и вновь ждёт сигнал в канале `ch`.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // создаём канал типа string, который даёт семантику сигналов;
    // небуферизованный канал гарантирует, что отправленный сигнал будет принят.
    ch := make(chan string)

    // количество горутин; неплохая отправная точка — runtime.NumCPU()
    // g := runtime.NumCPU()
    g := 3

    for e := 0; e < g; e++ {
        // для каждого сотрудника создаём отдельную горутину
        go func(emp int) {
            // сотрудник ждёт сигнал о том, что появилась работа;
            // все горутины блокируются на одном канале `ch` (получение).
            for p := range ch {
                fmt.Printf("employee %d : received signal : %s\n", emp, p)
            }

            // когда вся работа отправлена, менеджер уведомляет сотрудников,
            // закрывая канал; как только канал закрыт, for‑range завершается
            // и сотрудник выходит из цикла.
            fmt.Printf("employee %d : revieved shutdown signal\n", emp)
        }(e)
    }

    // объём работы
    const work = 10

    for w := 0; w < work; w++ {
        // когда работа готова, отправляем сигнал от менеджера сотрудникам;
        // отправитель (manager) уверен, что worker (employee) получил сигнал;
        // менеджеру не важно, какой именно сотрудник получил сигнал —
        // все способны выполнить задачу.
        ch <- "paper"

        fmt.Println("manager : sent signal :", w)
    }

    // когда вся работа отправлена, менеджер уведомляет всех закрытием канала;
    // небуферизованный канал гарантирует, что вся работа была доставлена.
    close(ch)
    fmt.Println("manager : sent shutdown signal")

    time.Sleep(time.Second)

}

```

### Результат (1‑й запуск)

```
go run main.go

employee 2 : recieved signal : paper
manager : sent signal : 0
manager : sent signal : 1
manager : sent signal : 2
manager : sent signal : 3
employee 1 : recieved signal : paper
employee 1 : recieved signal : paper
employee 2 : recieved signal : paper
manager : sent signal : 4
manager : sent signal : 5
manager : sent signal : 6
employee 1 : recieved signal : paper
employee 1 : recieved signal : paper
employee 0 : recieved signal : paper
employee 2 : recieved signal : paper
manager : sent signal : 7
manager : sent signal : 8
manager : sent signal : 9
manager : sent shutdown signal
employee 0 : recieved signal : paper
employee 0 : revieved shutdown signal
employee 2 : revieved shutdown signal
employee 1 : recieved signal : paper
employee 1 : revieved shutdown signal
```