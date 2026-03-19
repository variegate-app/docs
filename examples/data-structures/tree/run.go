package main

import "fmt"

type Node struct {
	val      int
	children []*Node
}

type Tree struct {
	root *Node
}

// O(n): для наглядности доступ по индексу в BFS-порядке.
func (t *Tree) Access(index int) (int, bool) {
	if t.root == nil {
		return 0, false
	}
	q := []*Node{t.root}
	for i := 0; len(q) > 0; i++ {
		cur := q[0]
		q = q[1:]
		if i == index {
			return cur.val, true
		}
		q = append(q, cur.children...)
	}
	return 0, false
}

// O(n): DFS по узлам.
func (t *Tree) Search(target int) bool {
	var dfs func(*Node) bool
	dfs = func(n *Node) bool {
		if n == nil {
			return false
		}
		if n.val == target {
			return true
		}
		for _, c := range n.children {
			if dfs(c) {
				return true
			}
		}
		return false
	}
	return dfs(t.root)
}

// O(1): вставка нового ребенка по указателю на родителя.
func (t *Tree) Insert(parent *Node, value int) *Node {
	n := &Node{val: value}
	if parent == nil {
		t.root = n
		return n
	}
	parent.children = append(parent.children, n)
	return n
}

// O(n): поиск и удаление поддерева.
func (t *Tree) Delete(target int) bool {
	if t.root == nil {
		return false
	}
	if t.root.val == target {
		t.root = nil
		return true
	}
	q := []*Node{t.root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i, c := range cur.children {
			if c.val == target {
				cur.children = append(cur.children[:i], cur.children[i+1:]...)
				return true
			}
			q = append(q, c)
		}
	}
	return false
}

func main() {
	t := &Tree{}
	r := t.Insert(nil, 1)
	t.Insert(r, 2)
	n3 := t.Insert(r, 3)
	t.Insert(n3, 4)
	fmt.Println("Search(4):", t.Search(4))
	v, _ := t.Access(2)
	fmt.Println("Access(2):", v)
	t.Delete(3)
	fmt.Println("Search(4) after Delete(3):", t.Search(4))
}
