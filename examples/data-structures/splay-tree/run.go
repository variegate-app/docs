package main

import "fmt"

type splayNode struct {
	val          int
	left, right  *splayNode
}

type SplayTree struct {
	root *splayNode
}

func rotateRight(x *splayNode) *splayNode {
	y := x.left
	x.left = y.right
	y.right = x
	return y
}

func rotateLeft(x *splayNode) *splayNode {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

func splay(root *splayNode, key int) *splayNode {
	if root == nil || root.val == key {
		return root
	}
	if key < root.val {
		if root.left == nil {
			return root
		}
		if key < root.left.val {
			root.left.left = splay(root.left.left, key)
			root = rotateRight(root)
		} else if key > root.left.val {
			root.left.right = splay(root.left.right, key)
			if root.left.right != nil {
				root.left = rotateLeft(root.left)
			}
		}
		if root.left == nil {
			return root
		}
		return rotateRight(root)
	}
	if root.right == nil {
		return root
	}
	if key > root.right.val {
		root.right.right = splay(root.right.right, key)
		root = rotateLeft(root)
	} else if key < root.right.val {
		root.right.left = splay(root.right.left, key)
		if root.right.left != nil {
			root.right = rotateRight(root.right)
		}
	}
	if root.right == nil {
		return root
	}
	return rotateLeft(root)
}

// O(log n) amortized: узел "подтягивается" к корню.
func (t *SplayTree) Access(target int) (int, bool) {
	t.root = splay(t.root, target)
	if t.root != nil && t.root.val == target {
		return t.root.val, true
	}
	return 0, false
}

// O(log n) amortized.
func (t *SplayTree) Search(target int) bool {
	_, ok := t.Access(target)
	return ok
}

// O(log n) amortized.
func (t *SplayTree) Insert(value int) {
	if t.root == nil {
		t.root = &splayNode{val: value}
		return
	}
	t.root = splay(t.root, value)
	if t.root.val == value {
		return
	}
	n := &splayNode{val: value}
	if value < t.root.val {
		n.right = t.root
		n.left = t.root.left
		t.root.left = nil
	} else {
		n.left = t.root
		n.right = t.root.right
		t.root.right = nil
	}
	t.root = n
}

// O(log n) amortized.
func (t *SplayTree) Delete(value int) bool {
	if t.root == nil {
		return false
	}
	t.root = splay(t.root, value)
	if t.root.val != value {
		return false
	}
	if t.root.left == nil {
		t.root = t.root.right
		return true
	}
	right := t.root.right
	t.root = splay(t.root.left, value)
	t.root.right = right
	return true
}

func main() {
	t := &SplayTree{}
	for _, v := range []int{10, 20, 30, 40, 50} {
		t.Insert(v)
	}
	fmt.Println("Search(30):", t.Search(30))
	v, ok := t.Access(40)
	fmt.Println("Access(40):", v, ok)
	t.Delete(20)
	fmt.Println("Search(20) after delete:", t.Search(20))
}
