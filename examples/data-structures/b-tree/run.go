package main

import (
	"fmt"
	"sort"
)

// Упрощенный B-Tree-подобный пример: данные в отсортированном блоке.
// Демонстрирует логарифмический поиск и вставку/удаление через позицию.
type BTreeDemo struct {
	keys []int
}

// O(log n): binary search по отсортированному блоку.
func (t *BTreeDemo) Access(index int) (int, bool) {
	if index < 0 || index >= len(t.keys) {
		return 0, false
	}
	return t.keys[index], true
}

// O(log n): бинарный поиск.
func (t *BTreeDemo) Search(target int) int {
	i := sort.SearchInts(t.keys, target)
	if i < len(t.keys) && t.keys[i] == target {
		return i
	}
	return -1
}

// O(log n) для поиска позиции + O(n) сдвиг в слайсе.
// В классическом B-Tree с узлами на диске работа остается O(log n).
func (t *BTreeDemo) Insert(value int) {
	i := sort.SearchInts(t.keys, value)
	t.keys = append(t.keys, 0)
	copy(t.keys[i+1:], t.keys[i:])
	t.keys[i] = value
}

// O(log n) поиск + O(n) сдвиг в слайсе (в B-Tree: O(log n) по узлам).
func (t *BTreeDemo) Delete(value int) bool {
	i := t.Search(value)
	if i < 0 {
		return false
	}
	t.keys = append(t.keys[:i], t.keys[i+1:]...)
	return true
}

func main() {
	t := &BTreeDemo{}
	t.Insert(10)
	t.Insert(5)
	t.Insert(20)
	t.Insert(15)
	fmt.Println("Search(15):", t.Search(15))
	v, _ := t.Access(1)
	fmt.Println("Access(1):", v)
	t.Delete(10)
	fmt.Println("After Delete(10):", t.keys)
}
