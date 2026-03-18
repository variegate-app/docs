# Команда (Command)

## [<<< ---](../../index.md)

## Назначение

Команда инкапсулирует запрос в объект, позволяя параметризовать действия, ставить их в очередь, логировать и поддерживать отмену.

Разделяет «кто вызывает» и «что именно делается».

## Когда использовать

- когда вам нужны операции, которые можно ставить в очередь или выполнять асинхронно;
- когда требуется поддержка undo/redo (или хотя бы логирования действий);
- когда хотите избежать больших условных блоков в коде контроллера.

## Пример реализации на Go

```go
package main

import "fmt"

// Receiver — знает, как делать реальное действие.
type Receiver struct{}

func (r Receiver) DoAction(arg string) { fmt.Println("do:", arg) }

// Command — оборачивает receiver+аргументы.
type Command interface {
	Execute()
}

type SimpleCommand struct {
	receiver Receiver
	arg      string
}

func (c SimpleCommand) Execute() { c.receiver.DoAction(c.arg) }

// Invoker — вызывает команды.
type Invoker struct {
	queue []Command
}

func (i *Invoker) Add(cmd Command) { i.queue = append(i.queue, cmd) }
func (i *Invoker) Run() {
	for _, cmd := range i.queue {
		cmd.Execute()
	}
}

func main() {
	inv := &Invoker{}
	r := Receiver{}
	inv.Add(SimpleCommand{receiver: r, arg: "task-1"})
	inv.Add(SimpleCommand{receiver: r, arg: "task-2"})
	inv.Run()
}
```

