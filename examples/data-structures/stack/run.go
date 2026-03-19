package main

import "fmt"

type Stack struct {
	data []int
}

// O(1): доступ к вершине.
func (s *Stack) Access() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}

// O(n): линейный поиск.
func (s *Stack) Search(target int) int {
	for i := len(s.data) - 1; i >= 0; i-- {
		if s.data[i] == target {
			return i
		}
	}
	return -1
}

// O(1) амортизированно.
func (s *Stack) Insert(value int) {
	s.data = append(s.data, value)
}

// O(1): pop с конца слайса.
func (s *Stack) Delete() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	i := len(s.data) - 1
	v := s.data[i]
	s.data = s.data[:i]
	return v, true
}

func main() {
	s := &Stack{}
	s.Insert(10)
	s.Insert(20)
	top, _ := s.Access()
	fmt.Println("Access(top):", top)
	fmt.Println("Search(10):", s.Search(10))
	pop, _ := s.Delete()
	fmt.Println("Delete(pop):", pop, "stack:", s.data)
}
