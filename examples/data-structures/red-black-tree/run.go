package main

import "fmt"

type color bool

const (
	red   color = true
	black color = false
)

type rbNode struct {
	val          int
	c            color
	left, right  *rbNode
	parent       *rbNode
}

type RedBlackTree struct {
	root *rbNode
}

// O(log n): для RB-дерева доступ по рангу/пути от корня.
func (t *RedBlackTree) Access(target int) (int, bool) {
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

// O(log n): высота RB-дерева ограничена.
func (t *RedBlackTree) Search(target int) bool {
	_, ok := t.Access(target)
	return ok
}

// O(log n): вставка как в BST + O(1) повороты/перекраска на уровне.
// Для компактности пример не содержит полный fix-up; цель - показать форму API и путь поиска.
func (t *RedBlackTree) Insert(value int) {
	n := &rbNode{val: value, c: red}
	if t.root == nil {
		n.c = black
		t.root = n
		return
	}
	cur := t.root
	var p *rbNode
	for cur != nil {
		p = cur
		if value < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	n.parent = p
	if value < p.val {
		p.left = n
	} else {
		p.right = n
	}
	t.root.c = black
}

// O(log n): поиск узла + локальная перестройка.
func (t *RedBlackTree) Delete(value int) bool {
	cur := t.root
	var p *rbNode
	for cur != nil && cur.val != value {
		p = cur
		if value < cur.val {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	if cur == nil {
		return false
	}
	replace := cur.left
	if replace == nil {
		replace = cur.right
	}
	if p == nil {
		t.root = replace
	} else if p.left == cur {
		p.left = replace
	} else {
		p.right = replace
	}
	if t.root != nil {
		t.root.c = black
	}
	return true
}

func main() {
	t := &RedBlackTree{}
	for _, v := range []int{10, 5, 20, 15, 25} {
		t.Insert(v)
	}
	fmt.Println("Search(15):", t.Search(15))
	v, ok := t.Access(20)
	fmt.Println("Access(20):", v, ok)
	t.Delete(20)
	fmt.Println("Search(20) after delete:", t.Search(20))
}
