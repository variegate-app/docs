package main

import "fmt"

type point struct {
	x, y int
}

type kdNode struct {
	p           point
	left, right *kdNode
}

type KDTree struct {
	root *kdNode
	pts  []point
}

func less(a, b point, depth int) bool {
	if depth%2 == 0 {
		return a.x < b.x
	}
	return a.y < b.y
}

// O(log n) average: путь по осям x/y.
func (t *KDTree) Access(target point) (point, bool) {
	cur := t.root
	depth := 0
	for cur != nil {
		if cur.p == target {
			return cur.p, true
		}
		if less(target, cur.p, depth) {
			cur = cur.left
		} else {
			cur = cur.right
		}
		depth++
	}
	return point{}, false
}

// O(log n) average.
func (t *KDTree) Search(target point) bool {
	_, ok := t.Access(target)
	return ok
}

func insertKD(n *kdNode, p point, depth int) *kdNode {
	if n == nil {
		return &kdNode{p: p}
	}
	if less(p, n.p, depth) {
		n.left = insertKD(n.left, p, depth+1)
	} else {
		n.right = insertKD(n.right, p, depth+1)
	}
	return n
}

// O(log n) average.
func (t *KDTree) Insert(p point) {
	t.root = insertKD(t.root, p, 0)
	t.pts = append(t.pts, p)
}

// O(n log n): удаление точки из списка + пересборка дерева.
func (t *KDTree) Delete(target point) bool {
	idx := -1
	for i, p := range t.pts {
		if p == target {
			idx = i
			break
		}
	}
	if idx < 0 {
		return false
	}
	t.pts = append(t.pts[:idx], t.pts[idx+1:]...)
	t.root = nil
	for _, p := range t.pts {
		t.root = insertKD(t.root, p, 0)
	}
	return true
}

func main() {
	t := &KDTree{}
	t.Insert(point{3, 6})
	t.Insert(point{17, 15})
	t.Insert(point{13, 15})
	t.Insert(point{6, 12})
	fmt.Println("Search({6,12}):", t.Search(point{6, 12}))
	p, ok := t.Access(point{13, 15})
	fmt.Println("Access({13,15}):", p, ok)
	fmt.Println("Delete({3,6}):", t.Delete(point{3, 6}))
	fmt.Println("Search({3,6}) after delete:", t.Search(point{3, 6}))
}
