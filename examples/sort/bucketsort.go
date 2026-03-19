package sort

/*
BucketSort

Временная сложность: O(n+k) в лучшем и среднем случаях. В худшем случае — O(n^2).
Пространственная сложность: O(n)

Применение:
  - Обработка данных с известным диапазоном значений:
    Если данные имеют ограниченный диапазон значений, Bucket Sort может быть эффективен,
    так как позволяет равномерно распределить элементы по корзинам.
  - Работа с числовыми данными:
    Особенно полезен для сортировки числовых данных, которые можно распределить по блокам на основе диапазонов значений.
  - Анализ и визуализация данных:
    Используется в аналитике данных для предварительной сортировки данных перед более детальным анализом.
  - Обработка изображений и сигналов:
    Применяется в обработке сигналов и изображений для сортировки значений пикселей или частотных компонентов по определенным диапазонам.
  - Образовательные цели:
    Часто используется в учебных курсах по алгоритмам для объяснения концепций сортировки и распределения данных.

Принцип работы: Этот алгоритм работает путем разделения элементов на фиксированное количество корзин,
каждая из которых содержит диапазон значений.
Затем элементы в каждой корзине сортируются с использованием другого алгоритма сортировки,
такого как сортировка вставками или быстрая сортировка.
После этого отсортированные корзины объединяются обратно.
*/
func BucketSort(arr []int) []int {
	maxVal := 0
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}

	buckets := make([][]int, maxVal+1)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	for _, val := range arr {
		buckets[val] = append(buckets[val], val)
	}

	result := make([]int, 0)
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}

func BucketSortCustom(arr []int) []int {
	buckets := make([][]int, len(arr))
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	for _, val := range arr {
		index := int(val * len(arr))
		buckets[index] = append(buckets[index], val)
	}

	for _, bucket := range buckets {
		for i := 1; i < len(bucket); i++ {
			j := i
			for j > 0 && bucket[j-1] > bucket[j] {
				bucket[j-1], bucket[j] = bucket[j], bucket[j-1]
				j--
			}
		}
	}

	result := make([]int, 0)
	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}
