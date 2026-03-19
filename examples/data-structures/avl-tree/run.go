package main

import "fmt"

type avlNode struct {
	val          int
	height       int
	left, right  *avlNode
}

type AVLTree struct {
	root *avlNode
}

func h(n *avlNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

func upd(n *avlNode) {
	if n == nil {
		return
	}
	hl, hr := h(n.left), h(n.right)
	if hl > hr {
		n.height = hl + 1
	} else {
		n.height = hr + 1
	}
}

func bf(n *avlNode) int { return h(n.left) - h(n.right) }

func rotR(y *avlNode) *avlNode {
	x := y.left
	t2 := x.right
	x.right = y
	y.left = t2
	upd(y)
	upd(x)
	return x
}

func rotL(x *avlNode) *avlNode {
	y := x.right
	t2 := y.left
	y.left = x
	x.right = t2
	upd(x)
	upd(y)
	return y
}

func rebalance(n *avlNode) *avlNode {
	upd(n)
	b := bf(n)
	if b > 1 {
		if bf(n.left) < 0 {
			n.left = rotL(n.left)
		}
		return rotR(n)
	}
	if b < -1 {
		if bf(n.right) > 0 {
			n.right = rotR(n.right)
		}
		return rotL(n)
	}
	return n
}

// O(log n): путь от корня ограничен балансировкой.
func (t *AVLTree) Access(target int) (int, bool) {
	cur := t.root
	for cur != nil {
		if target == cur.val {
			return cur.val, true
		}
		if target < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return 0, false
}

// O(log n).
func (t *AVLTree) Search(target int) bool {
	_, ok := t.Access(target)
	return ok
}

func insert(n *avlNode, v int) *avlNode {
	if n == nil {
		return &avlNode{val: v, height: 1}
	}
	if v < n.val {
		n.left = insert(n.left, v)
	} else if v > n.val {
		n.right = insert(n.right, v)
	}
	return rebalance(n)
}

// O(log n): после вставки максимум O(log n) поворотов.
func (t *AVLTree) Insert(v int) { t.root = insert(t.root, v) }

func minAVL(n *avlNode) *avlNode {
	for n.left != nil {
		n = n.left
	}
	return n
}

func remove(n *avlNode, v int) *avlNode {
	if n == nil {
		return nil
	}
	if v < n.val {
		n.left = remove(n.left, v)
	} else if v > n.val {
		n.right = remove(n.right, v)
	} else {
		if n.left == nil || n.right == nil {
			if n.left != nil {
				return n.left
			}
			return n.right
		}
		s := minAVL(n.right)
		n.val = s.val
		n.right = remove(n.right, s.val)
	}
	return rebalance(n)
}

// O(log n).
func (t *AVLTree) Delete(v int) { t.root = remove(t.root, v) }

func main() {
	t := &AVLTree{}
	for _, v := range []int{30, 20, 40, 10, 25, 35, 50} {
		t.Insert(v)
	}
	fmt.Println("Search(25):", t.Search(25))
	v, ok := t.Access(35)
	fmt.Println("Access(35):", v, ok)
	t.Delete(20)
	fmt.Println("Search(20) after delete:", t.Search(20))
}
