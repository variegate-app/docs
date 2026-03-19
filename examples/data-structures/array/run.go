package main

import "fmt"

type Array struct {
	data []int
}

// O(1): прямой доступ по индексу.
func (a *Array) Access(index int) (int, bool) {
	if index < 0 || index >= len(a.data) {
		return 0, false
	}
	return a.data[index], true
}

// O(n): линейный поиск.
func (a *Array) Search(target int) int {
	for i, v := range a.data {
		if v == target {
			return i
		}
	}
	return -1
}

// O(n): нужно сдвигать хвост вправо.
func (a *Array) Insert(index, value int) bool {
	if index < 0 || index > len(a.data) {
		return false
	}
	a.data = append(a.data, 0)
	copy(a.data[index+1:], a.data[index:])
	a.data[index] = value
	return true
}

// O(n): нужно сдвигать хвост влево.
func (a *Array) Delete(index int) bool {
	if index < 0 || index >= len(a.data) {
		return false
	}
	copy(a.data[index:], a.data[index+1:])
	a.data = a.data[:len(a.data)-1]
	return true
}

func main() {
	a := &Array{data: []int{10, 20, 30, 40}}
	v, _ := a.Access(2)
	fmt.Println("Access(2):", v)
	fmt.Println("Search(30):", a.Search(30))
	a.Insert(1, 15)
	fmt.Println("After Insert(1,15):", a.data)
	a.Delete(3)
	fmt.Println("After Delete(3):", a.data)
}
