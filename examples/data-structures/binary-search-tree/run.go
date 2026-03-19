package main

import "fmt"

type bstNode struct {
	val         int
	left, right *bstNode
}

type BST struct {
	root *bstNode
}

// O(log n) average: путь от корня к i-му элементу в inorder.
func (t *BST) Access(index int) (int, bool) {
	stack := []*bstNode{}
	cur := t.root
	i := 0
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if i == index {
			return cur.val, true
		}
		i++
		cur = cur.right
	}
	return 0, false
}

// O(log n) average, O(n) worst.
func (t *BST) Search(target int) bool {
	cur := t.root
	for cur != nil {
		if target == cur.val {
			return true
		}
		if target < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return false
}

// O(log n) average.
func (t *BST) Insert(value int) {
	if t.root == nil {
		t.root = &bstNode{val: value}
		return
	}
	cur := t.root
	for {
		if value < cur.val {
			if cur.left == nil {
				cur.left = &bstNode{val: value}
				return
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &bstNode{val: value}
				return
			}
			cur = cur.right
		}
	}
}

func deleteNode(root *bstNode, key int) *bstNode {
	if root == nil {
		return nil
	}
	if key < root.val {
		root.left = deleteNode(root.left, key)
		return root
	}
	if key > root.val {
		root.right = deleteNode(root.right, key)
		return root
	}
	if root.left == nil {
		return root.right
	}
	if root.right == nil {
		return root.left
	}
	s := root.right
	for s.left != nil {
		s = s.left
	}
	root.val = s.val
	root.right = deleteNode(root.right, s.val)
	return root
}

// O(log n) average.
func (t *BST) Delete(value int) {
	t.root = deleteNode(t.root, value)
}

func main() {
	t := &BST{}
	for _, v := range []int{8, 3, 10, 1, 6, 14} {
		t.Insert(v)
	}
	fmt.Println("Search(6):", t.Search(6))
	v, _ := t.Access(2)
	fmt.Println("Access(2):", v)
	t.Delete(3)
	fmt.Println("Search(3) after delete:", t.Search(3))
}
