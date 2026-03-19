package main

import "fmt"

type entry struct {
	key   string
	value int
}

type HashTable struct {
	buckets [][]entry
}

func NewHashTable(size int) *HashTable {
	return &HashTable{buckets: make([][]entry, size)}
}

func (h *HashTable) hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % len(h.buckets)
}

// N/A: у hash table нет доступа по числовому индексу.
func (h *HashTable) Access(_ int) (int, bool) {
	return 0, false
}

// O(1)* среднее, O(n) худшее при коллизиях.
func (h *HashTable) Search(key string) (int, bool) {
	b := h.buckets[h.hash(key)]
	for _, e := range b {
		if e.key == key {
			return e.value, true
		}
	}
	return 0, false
}

// O(1)* среднее, O(n) худшее.
func (h *HashTable) Insert(key string, value int) {
	i := h.hash(key)
	for j := range h.buckets[i] {
		if h.buckets[i][j].key == key {
			h.buckets[i][j].value = value
			return
		}
	}
	h.buckets[i] = append(h.buckets[i], entry{key: key, value: value})
}

// O(1)* среднее, O(n) худшее.
func (h *HashTable) Delete(key string) bool {
	i := h.hash(key)
	b := h.buckets[i]
	for j, e := range b {
		if e.key == key {
			h.buckets[i] = append(b[:j], b[j+1:]...)
			return true
		}
	}
	return false
}

func main() {
	h := NewHashTable(8)
	h.Insert("go", 1)
	h.Insert("rust", 2)
	v, _ := h.Search("go")
	fmt.Println("Search(go):", v)
	h.Delete("rust")
	_, ok := h.Search("rust")
	fmt.Println("Search(rust):", ok)
}
