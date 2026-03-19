package main

import "fmt"

type MinHeap struct {
	data []int
}

// O(1): доступ к минимальному элементу в корне.
func (h *MinHeap) Access() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

// O(n): heap не поддерживает эффективный произвольный поиск.
func (h *MinHeap) Search(target int) int {
	for i, v := range h.data {
		if v == target {
			return i
		}
	}
	return -1
}

// O(log n): подъем элемента вверх (sift up).
func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	i := len(h.data) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h.data[p] <= h.data[i] {
			break
		}
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

// O(log n): удаление root + sift down.
func (h *MinHeap) Delete() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	min := h.data[0]
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) == 0 {
		return min, true
	}
	h.data[0] = last
	i := 0
	for {
		l, r := 2*i+1, 2*i+2
		small := i
		if l < len(h.data) && h.data[l] < h.data[small] {
			small = l
		}
		if r < len(h.data) && h.data[r] < h.data[small] {
			small = r
		}
		if small == i {
			break
		}
		h.data[i], h.data[small] = h.data[small], h.data[i]
		i = small
	}
	return min, true
}

func main() {
	h := &MinHeap{}
	h.Insert(7)
	h.Insert(3)
	h.Insert(9)
	h.Insert(1)
	min, _ := h.Access()
	fmt.Println("Access(min):", min)
	fmt.Println("Search(9):", h.Search(9))
	del, _ := h.Delete()
	fmt.Println("Delete(min):", del, "heap:", h.data)
}
