package sort

type tree struct {
	value       int
	left, right *tree
}

/*
TreeSort

Временная сложность: в среднем O(n log n) , в худшем случае O(n^2)
Пространственная сложность: O(n)

Преимущество: применяется, когда данные поступают из потока (например, файла, сокета или консоли)
Недостаток:
- Неэффективен для небольших наборов данных из-за затрат на построение и обход дерева.
- Требует дополнительной памяти для структуры дерева, что приводит к высокой пространственной сложности.
- Производительность сильно зависит от баланса дерева: несбалансированное дерево может снизить эффективность сортировки до O(n^2)

Принцип работы: Заключается в построении двоичного дерева поиска по ключам массива (списка),

	с последующей сборкой результирующего массива путём обхода узлов построенного дерева,
	в необходимом порядке следования ключей
*/
func TreeSort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = root.add(v)
	}
	return root.appendValues(values[:0])
}

func (t *tree) appendValues(values []int) []int {
	if t != nil {
		values = t.left.appendValues(values)
		values = append(values, t.value)
		values = t.right.appendValues(values)
	}
	return values
}

func (t *tree) add(value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = t.left.add(value)
	} else {
		t.right = t.right.add(value)
	}
	return t
}
