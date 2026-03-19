package main

import "fmt"

type Queue struct {
	data []int
}

// O(1): доступ к голове очереди.
func (q *Queue) Access() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	return q.data[0], true
}

// O(n): надо просмотреть элементы.
func (q *Queue) Search(target int) int {
	for i, v := range q.data {
		if v == target {
			return i
		}
	}
	return -1
}

// O(1) амортизированно: append в конец.
func (q *Queue) Insert(value int) {
	q.data = append(q.data, value)
}

// O(1): удаление головы через реслайс.
func (q *Queue) Delete() (int, bool) {
	if len(q.data) == 0 {
		return 0, false
	}
	v := q.data[0]
	q.data[0] = 0
	q.data = q.data[1:]
	return v, true
}

func main() {
	q := &Queue{}
	q.Insert(1)
	q.Insert(2)
	q.Insert(3)
	v, _ := q.Access()
	fmt.Println("Access(head):", v)
	fmt.Println("Search(3):", q.Search(3))
	d, _ := q.Delete()
	fmt.Println("Delete(head):", d, "queue:", q.data)
}
