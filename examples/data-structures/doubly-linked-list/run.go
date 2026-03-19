package main

import "fmt"

type dnode struct {
	val  int
	prev *dnode
	next *dnode
}

type DoublyLinkedList struct {
	head *dnode
	tail *dnode
}

// O(n): доступ по индексу.
func (l *DoublyLinkedList) Access(index int) (int, bool) {
	cur := l.head
	for i := 0; cur != nil; i++ {
		if i == index {
			return cur.val, true
		}
		cur = cur.next
	}
	return 0, false
}

// O(n): линейный поиск.
func (l *DoublyLinkedList) Search(target int) int {
	cur := l.head
	for i := 0; cur != nil; i++ {
		if cur.val == target {
			return i
		}
		cur = cur.next
	}
	return -1
}

// O(1): вставка в хвост при наличии tail.
func (l *DoublyLinkedList) Insert(value int) {
	n := &dnode{val: value}
	if l.tail == nil {
		l.head, l.tail = n, n
		return
	}
	l.tail.next = n
	n.prev = l.tail
	l.tail = n
}

// O(1): удаление хвоста.
func (l *DoublyLinkedList) Delete() (int, bool) {
	if l.tail == nil {
		return 0, false
	}
	v := l.tail.val
	if l.tail.prev == nil {
		l.head, l.tail = nil, nil
		return v, true
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	return v, true
}

func main() {
	l := &DoublyLinkedList{}
	l.Insert(10)
	l.Insert(20)
	l.Insert(30)
	v, _ := l.Access(1)
	fmt.Println("Access(1):", v)
	fmt.Println("Search(30):", l.Search(30))
	d, _ := l.Delete()
	fmt.Println("Delete(tail):", d)
}
