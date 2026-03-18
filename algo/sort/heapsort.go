package sort

/*
HeapSort

Временная сложность: O(n log n) в худшем, среднем и лучшем случаях.
Пространственная сложность: O(1).

Преимущество: Очень быстрая сортировка, подходит для больших данных.
Недостаток: Не стабилен (порядок равных элементов может измениться).

Принцип работы: Строится двоичная куча, затем из неё поочередно извлекаются элементы,
и в итоге получается отсортированный массив.

Шаги:
- Строим кучу из массива.
- Извлекаем элементы из кучи поочередно и перестраиваем кучу после каждого извлечения.
*/
func HeapSort(numbers []int) {
	n := len(numbers)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(numbers, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		numbers[0], numbers[i] = numbers[i], numbers[0]
		heapify(numbers, i, 0)
	}
}

func heapify(numbers []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && numbers[left] > numbers[largest] {
		largest = left
	}
	if right < n && numbers[right] > numbers[largest] {
		largest = right
	}

	if largest != i {
		numbers[i], numbers[largest] = numbers[largest], numbers[i]
		heapify(numbers, n, largest)
	}
}
