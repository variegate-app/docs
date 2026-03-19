package main

import "fmt"

type node struct {
	val  int
	next *node
}

type SinglyLinkedList struct {
	head *node
}

// O(n): доступ по индексу требует прохода по ссылкам.
func (l *SinglyLinkedList) Access(index int) (int, bool) {
	cur := l.head
	for i := 0; cur != nil; i++ {
		if i == index {
			return cur.val, true
		}
		cur = cur.next
	}
	return 0, false
}

// O(n): линейный проход.
func (l *SinglyLinkedList) Search(target int) int {
	cur := l.head
	for i := 0; cur != nil; i++ {
		if cur.val == target {
			return i
		}
		cur = cur.next
	}
	return -1
}

// O(1): вставка в голову.
func (l *SinglyLinkedList) Insert(value int) {
	l.head = &node{val: value, next: l.head}
}

// O(1): удаление головы.
func (l *SinglyLinkedList) Delete() (int, bool) {
	if l.head == nil {
		return 0, false
	}
	v := l.head.val
	l.head = l.head.next
	return v, true
}

func main() {
	l := &SinglyLinkedList{}
	l.Insert(3)
	l.Insert(2)
	l.Insert(1)
	v, _ := l.Access(1)
	fmt.Println("Access(1):", v)
	fmt.Println("Search(3):", l.Search(3))
	d, _ := l.Delete()
	fmt.Println("Delete(head):", d)
}
