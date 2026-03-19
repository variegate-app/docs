package main

import "fmt"

type btNode struct {
	val         int
	left, right *btNode
}

type BinaryTree struct {
	root *btNode
}

// O(n): доступ к i-му элементу в BFS-порядке.
func (t *BinaryTree) Access(index int) (int, bool) {
	if t.root == nil {
		return 0, false
	}
	q := []*btNode{t.root}
	for i := 0; len(q) > 0; i++ {
		cur := q[0]
		q = q[1:]
		if i == index {
			return cur.val, true
		}
		if cur.left != nil {
			q = append(q, cur.left)
		}
		if cur.right != nil {
			q = append(q, cur.right)
		}
	}
	return 0, false
}

// O(n): без свойства порядка нужен полный обход.
func (t *BinaryTree) Search(target int) bool {
	var dfs func(*btNode) bool
	dfs = func(n *btNode) bool {
		if n == nil {
			return false
		}
		if n.val == target {
			return true
		}
		return dfs(n.left) || dfs(n.right)
	}
	return dfs(t.root)
}

// O(n): для демонстрации вставляем в первую свободную позицию (BFS).
func (t *BinaryTree) Insert(value int) {
	n := &btNode{val: value}
	if t.root == nil {
		t.root = n
		return
	}
	q := []*btNode{t.root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.left == nil {
			cur.left = n
			return
		}
		if cur.right == nil {
			cur.right = n
			return
		}
		q = append(q, cur.left, cur.right)
	}
}

// O(n): ищем узел и перестраиваем ссылку.
func (t *BinaryTree) Delete(target int) bool {
	if t.root == nil {
		return false
	}
	if t.root.val == target {
		t.root = nil
		return true
	}
	q := []*btNode{t.root}
	var targetNode, last, parentOfLast *btNode
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.val == target {
			targetNode = cur
		}
		if cur.left != nil {
			parentOfLast = cur
			last = cur.left
			q = append(q, cur.left)
		}
		if cur.right != nil {
			parentOfLast = cur
			last = cur.right
			q = append(q, cur.right)
		}
	}
	if targetNode == nil || last == nil {
		return false
	}
	targetNode.val = last.val
	if parentOfLast.left == last {
		parentOfLast.left = nil
	} else {
		parentOfLast.right = nil
	}
	return true
}

func main() {
	t := &BinaryTree{}
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	fmt.Println("Search(4):", t.Search(4))
	v, _ := t.Access(2)
	fmt.Println("Access(2):", v)
	t.Delete(2)
	fmt.Println("Search(2) after delete:", t.Search(2))
}
