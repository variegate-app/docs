package main

import (
	"fmt"
	"math/rand"
)

const maxLevel = 4

type skipNode struct {
	key  int
	next [maxLevel]*skipNode
}

type SkipList struct {
	head  *skipNode
	level int
}

func NewSkipList() *SkipList {
	return &SkipList{head: &skipNode{}, level: 1}
}

func randomLevel() int {
	lvl := 1
	for lvl < maxLevel && rand.Intn(2) == 0 {
		lvl++
	}
	return lvl
}

// O(log n) average: спуск по уровням.
func (s *SkipList) Access(index int) (int, bool) {
	cur := s.head.next[0]
	for i := 0; cur != nil; i++ {
		if i == index {
			return cur.key, true
		}
		cur = cur.next[0]
	}
	return 0, false
}

// O(log n) average.
func (s *SkipList) Search(key int) bool {
	cur := s.head
	for lvl := s.level - 1; lvl >= 0; lvl-- {
		for cur.next[lvl] != nil && cur.next[lvl].key < key {
			cur = cur.next[lvl]
		}
	}
	cur = cur.next[0]
	return cur != nil && cur.key == key
}

// O(log n) average.
func (s *SkipList) Insert(key int) {
	update := [maxLevel]*skipNode{}
	cur := s.head
	for lvl := s.level - 1; lvl >= 0; lvl-- {
		for cur.next[lvl] != nil && cur.next[lvl].key < key {
			cur = cur.next[lvl]
		}
		update[lvl] = cur
	}
	lvl := randomLevel()
	if lvl > s.level {
		for i := s.level; i < lvl; i++ {
			update[i] = s.head
		}
		s.level = lvl
	}
	n := &skipNode{key: key}
	for i := 0; i < lvl; i++ {
		n.next[i] = update[i].next[i]
		update[i].next[i] = n
	}
}

// O(log n) average.
func (s *SkipList) Delete(key int) bool {
	update := [maxLevel]*skipNode{}
	cur := s.head
	for lvl := s.level - 1; lvl >= 0; lvl-- {
		for cur.next[lvl] != nil && cur.next[lvl].key < key {
			cur = cur.next[lvl]
		}
		update[lvl] = cur
	}
	target := cur.next[0]
	if target == nil || target.key != key {
		return false
	}
	for i := 0; i < s.level; i++ {
		if update[i].next[i] != target {
			continue
		}
		update[i].next[i] = target.next[i]
	}
	return true
}

func main() {
	s := NewSkipList()
	s.Insert(10)
	s.Insert(20)
	s.Insert(15)
	fmt.Println("Search(15):", s.Search(15))
	v, _ := s.Access(1)
	fmt.Println("Access(1):", v)
	s.Delete(15)
	fmt.Println("Search(15) after delete:", s.Search(15))
}
