# Хранитель состояний (Memento)

## [<<< ---](../../index.md)

## Назначение

Memento сохраняет внутреннее состояние объекта так, чтобы позже можно было восстановить объект в прежний момент.

При этом само состояние хранится отдельно от объекта-«владельца».

## Когда использовать

- когда нужна возможность отката (undo) или возврата к ранее известному состоянию;
- когда вы хотите ограничить доступ к состоянию: клиенты не должны менять его напрямую;
- когда состояние нужно сериализовать/логировать как «снимок».

## Пример реализации на Go

```go
package main

import "fmt"

// Memento хранит снапшот состояния.
type Memento struct {
	state int
}

// Origin — объект, состояние которого можно сохранять и восстанавливать.
type Origin struct {
	state int
}

func (o *Origin) SetState(v int) { o.state = v }
func (o *Origin) GetState() int  { return o.state }

func (o *Origin) CreateMemento() Memento {
	return Memento{state: o.state}
}

func (o *Origin) Restore(m Memento) {
	o.state = m.state
}

func main() {
	origin := &Origin{state: 1}
	snap := origin.CreateMemento()

	origin.SetState(42)
	fmt.Println("after change:", origin.GetState())

	origin.Restore(snap)
	fmt.Println("after restore:", origin.GetState())
}
```

