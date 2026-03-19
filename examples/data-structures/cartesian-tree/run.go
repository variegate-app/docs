package main

import "fmt"

type cnode struct {
	val         int
	left, right *cnode
}

type CartesianTree struct {
	root *cnode
	data []int
}

func buildCartesian(arr []int) *cnode {
	if len(arr) == 0 {
		return nil
	}
	minI := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[minI] {
			minI = i
		}
	}
	return &cnode{
		val:   arr[minI],
		left:  buildCartesian(arr[:minI]),
		right: buildCartesian(arr[minI+1:]),
	}
}

// O(n): доступ по индексу в исходном массиве.
func (t *CartesianTree) Access(index int) (int, bool) {
	if index < 0 || index >= len(t.data) {
		return 0, false
	}
	return t.data[index], true
}

// O(n): линейный поиск по массиву.
func (t *CartesianTree) Search(target int) int {
	for i, v := range t.data {
		if v == target {
			return i
		}
	}
	return -1
}

// O(n): для наглядности пересобираем дерево после вставки.
func (t *CartesianTree) Insert(value int) {
	t.data = append(t.data, value)
	t.root = buildCartesian(t.data)
}

// O(n): удаляем из массива + пересборка дерева.
func (t *CartesianTree) Delete(value int) bool {
	i := t.Search(value)
	if i < 0 {
		return false
	}
	t.data = append(t.data[:i], t.data[i+1:]...)
	t.root = buildCartesian(t.data)
	return true
}

func main() {
	t := &CartesianTree{}
	t.Insert(5)
	t.Insert(2)
	t.Insert(8)
	t.Insert(1)
	v, _ := t.Access(2)
	fmt.Println("Access(2):", v)
	fmt.Println("Search(8):", t.Search(8))
	t.Delete(2)
	fmt.Println("Search(2) after delete:", t.Search(2))
}
