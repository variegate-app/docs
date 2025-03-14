package sort

/*
SelectionSort

Временная сложность: O(n^2) во всех случаях.
Пространственная сложность: O(1).

Преимущество: Простота реализации.
Недостаток: Не эффективен для больших массивов.

Принцип работы: В этом алгоритме на каждом шаге мы находим минимальный элемент в неотсортированной части массива и меняем его местами с первым элементом в этой части. Этот процесс повторяется, пока весь массив не будет отсортирован.
Шаги:
- Находим минимальный элемент в неотсортированной части массива.
- Меняем его местами с первым элементом неотсортированной части.
- Повторяем процесс для оставшейся части массива.
*/
func SelectionSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if numbers[j] < numbers[minIndex] {
				minIndex = j
			}
		}
		numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
	}
}
