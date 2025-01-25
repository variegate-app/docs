# Основные типы алгоритмов:
## [<<< ---](../README.md)

- **Линейные** 
    Это самый простой тип алгоритма: действия идут друг за другом, каждое начинается после того, как закончится предыдущее. Они не переставляются местами, не повторяются, выполняются при любых условиях.
- **Ветвящиеся**
    В этом типе алгоритма появляется ветвление: какие-то действия выполняются, только если верны некоторые условия. Например, если число меньше нуля, то его нужно удалить из структуры данных. Условий может быть несколько, они могут комбинироваться друг с другом.
- **Циклические** 
    Такие алгоритмы выполняются в цикле: когда какой-то блок действий заканчивается, эти действия начинаются снова и повторяются некоторое количество раз. Цикл может включать в себя одно действие или последовательность, а количество повторений может быть фиксированным или зависеть от условия. В некоторых случаях цикл может быть бесконечным.
- **Рекурсивные** 
    Рекурсия — это явление, когда какой-то алгоритм вызывает сам себя, но с другими входными данными. Известный пример рекурсивного алгоритма — расчёт чисел Фибоначчи.
- **Вероятностные**
    Работа алгоритма зависит не только от входных данных, но и от случайных величин. К ним, например, относятся известные алгоритмы Лас-Вегас и Монте-Карло.

# Сортировка

| Алгоритм | Time Best| Time Middle | Time Worst | Space Worst|
|---|---|---|---|---|
| Quick sort |<red>O(n log n)</red>|<red>O(n log n)</red>|<red>O(n)</red>|<green>O(n*log(n))</green>|
| Merge sort |<red>O(n log n)</red>|<red>O(n log n)</red>|<red>O(n log n)</red>| <blue>O(n) `#0969DA`|
| Tim sort |<blue>O(n)<blue>|<red>O(n log n)</red>|<red>O(n log n)</red>|<blue>O(1) `#0969DA`|
| Heap sort |<red>O(n log n)</red>|<red>O(n log n)</red>|<red>O(n log n)</red>|<green>O(1)</green>|
| Bubble sort |<blue>O(n) `#0969DA`|<red>O(n<sup>2</sup>)</red>|<red>O(n<sup>2</sup>)</red>|<green>O(1)</green>|
| Insertion Sort |<blue>O(n) `#0969DA`|<red>O(n<sup>2</sup>)</red>|<red>O(n<sup>2</sup>)</red>|<green>O(1)</green>|
| Selection Sort |<green>O(n<sup>2</sup>)</green>|<red>O(n<sup>2</sup>)</red>|<red>O(n<sup>2</sup>)</red>|<green>O(1)</green>|
| Tree Sort |<red>O(n log n)</red>|<red>O(n log(n) )</red>|<red>O(n<sup>2</sup>) * T</red>|<blue>O(n) `#0969DA`|
| Shell Sort |<red>O(n log n)</red>|<red>O(n (log n)<sup>2</sup>)</red>|<red>O(n (log n)<sup>2</sup>)</red>|<green>O(1)</green>|
| Bucket Sort |<green>O(n + K)</green>|<green>O(n + K)</green>|<red>O(n<sup>2</sup>)</red>|<blue>O(n) `#0969DA`|
| Radix Sort |<green>O( nK )</green>|<green>O( nK )</green>|<green>O( nK )</green>|<blue>O( n + K ) `#0969DA`|
| Counting Sort  |<green>O(n + K)</green>|<green>O(n + K)</green>|<green>O(n + K)</green>| <blue>O(K) `#0969DA`|
| Cube Sort  |<blue>O(n) `#0969DA`|<red>O(n log n)</red>|<red>O(n log n)</red>|<blue>O(n) `#0969DA`|
---
<details><summary>Quick sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/quicksort.go#L3-L38
</details>

<details><summary>Merge sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/mergesort.go#L3-L49
</details>

<details><summary>Tim sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/timsort.go#L3-L72
</details>

<details><summary>Heap sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/heapsort.go#L3-L48
</details>

<details><summary>Bubble sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/bubblesort.go#L3-L33
</details>

<details><summary>Insertion Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/insertionsort.go#L3-L28
</details>

<details><summary>Selection Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/selectionsort.go#L3-L29
</details>

<details><summary>Tree Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/treesort.go#L3-L54
</details>

<details><summary>Shell Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/shellsort.go#L3-L44
</details>

<details><summary>Bucket Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/bucketsort.go#L3-L80
</details>

<details><summary>Radix Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/radixsort.go#L3-L74
</details>

<details><summary>Counting Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/countingsort.go#L3-L59
</details>

<details><summary>Cube Sort</summary>
https://github.com/variegate-app/docs/blob/51261f42242e5a83ed3252ac190f2d53cbce847f/gomod/algo/sort/cubesort.go#L3-L46
</details>