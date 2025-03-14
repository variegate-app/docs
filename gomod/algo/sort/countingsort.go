package sort

/*
CountingSort

Временная сложность: O(n + k) во всех случаях, где k - диапазон значений ключа
Пространственная сложность: O(n + k).

Преимущество: Этот алгоритм особенно полезен в ситуациях, где диапазон значений фиксирован и элементы могут повторяться,
что делает его быстрым и эффективным для больших массивов.
Полезен, когда диапазон значений известен и намного меньше размера массива.
Это позволяет избежать сравнений, что делает его быстрым для больших массивов.

Принцип работы: Строится двоичная куча, затем из неё поочередно извлекаются элементы,
и в итоге получается отсортированный массив.

Часто используется как подпрограмма в более сложных алгоритмах сортировки, таких как Radix Sort,
для повышения эффективности обработки больших ключей.
*/
func CountingSort(arr []int) []int {
	var max = arr[0]

	var i = 1
	for i < len(arr) {
		if arr[i] > max {
			max = arr[i]
		}

		i++
	}

	var indices = make([]int, max+1)

	var j = 0
	for j < len(arr) {
		indices[arr[j]]++

		j++
	}

	var k = 1
	for k < len(indices) {
		indices[k] += indices[k-1]

		k++
	}

	var result = make([]int, len(arr))

	var m = 0
	for m < len(arr) {
		result[indices[arr[m]]-1] = arr[m]
		indices[arr[m]]--

		m++
	}

	return result
}
