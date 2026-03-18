# Layer Supertype

## [<<< ---](../../index.md)

Тип, выполняющий роль суперкласса для всех классов своего слоя. Довольно часто одни и те же методы дублируются во всех объектах слоя. Чтобы избежать повторений, все общее поведение можно вынести в Layer Supertype.

Типичным примером может стать `IdableEntity`, к примеру у вас есть несколько доменных классов и вы создали `IdableEntity` класс, в который поместили проперти айди, который общий для всех доменных объектов.

### Пример реализации на Go (Layer Supertype)

```go
package main

import "fmt"

// Layer Supertype задаёт общий “контракт” и общее поведение для сущностей слоя.
type IdableEntity struct {
	ID int64
}

func (e IdableEntity) GetID() int64 { return e.ID }

// Доменные объекты используют общий суперкласс (через embedding).
type User struct {
	IdableEntity
	Email string
}

type Order struct {
	IdableEntity
	Total int
}

func main() {
	u := User{IdableEntity: IdableEntity{ID: 1}, Email: "a@b.com"}
	o := Order{IdableEntity: IdableEntity{ID: 2}, Total: 100}

	fmt.Println(u.GetID(), o.GetID())
}
```