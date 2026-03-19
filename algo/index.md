## [<<< ---](../README.md)

# Основные типы алгоритмов:

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

# задачи

# Сортировка

| Алгоритм | Time Best| Time Middle | Time Worst | Space Worst|
|---|---|---|---|---|
| Quick sort |🔴 O(n log n)|🔴 O(n log n)|🔴 O(n)|🟢 O(n*log(n))|
| Merge sort |🔴 O(n log n)|🔴 O(n log n)|🔴 O(n log n)| 🔵 O(n)|
| Tim sort |🔵 O(n)|🔴 O(n log n)|🔴 O(n log n)|🔵 O(1)|
| Heap sort |🔴 O(n log n)|🔴 O(n log n)|🔴 O(n log n)|🟢 O(1)|
| Bubble sort |🔵 O(n)|🔴 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🟢 O(1)|
| Insertion Sort |🔵 O(n)|🔴 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🟢 O(1)|
| Selection Sort |🟢 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🟢 O(1)|
| Tree Sort |🔴 O(n log n)|🔴 O(n log(n) )|🔴 O(n<sup>2</sup>) * T|🔵 O(n)|
| Shell Sort |🔴 O(n log n)|🔴 O(n (log n)<sup>2</sup>)|🔴 O(n (log n)<sup>2</sup>)|🟢 O(1)|
| Bucket Sort |🟢 O(n + K)|🟢 O(n + K)|🔴 O(n<sup>2</sup>)|🔵 O(n)|
| Radix Sort |🟢 O( nK )|🟢 O( nK )|🟢 O( nK )|🔵 O( n + K )|
| Counting Sort  |🟢 O(n + K)|🟢 O(n + K)|🟢 O(n + K)| 🔵 O(K)|
| Cube Sort  |🔵 O(n)|🔴 O(n log n)|🔴 O(n log n)|🔵 O(n)|
---
<details><summary>Quick sort</summary>

![Quick](./assets/Quicksort.gif)
```golang:sort/quicksort.go [3-38]

```
</details>

<details><summary>Merge sort</summary>

![Merge](./assets/mergesort.gif)

./sort/mergesort.go#L3-L49
</details>

<details><summary>Tim sort</summary>
./sort/timsort.go#L3-L72
</details>

<details><summary>Heap sort</summary>

![Heap](./assets/heapsort.gif)

./sort/heapsort.go#L3-L48
</details>

<details><summary>Bubble sort</summary>

![Bubble](./assets/bubblesort.gif)

./sort/bubblesort.go#L3-L33
</details>

<details><summary>Insertion Sort</summary>

![Insert](./assets/insertionsort.gif) 

./sort/insertionsort.go#L3-L28
</details>

<details><summary>Selection Sort</summary>
./sort/selectionsort.go#L3-L29
</details>

<details><summary>Tree Sort</summary>
./sort/treesort.go#L3-L54
</details>

<details><summary>Shell Sort</summary>
./sort/shellsort.go#L3-L44
</details>

<details><summary>Bucket Sort</summary>
./sort/bucketsort.go#L3-L80
</details>

<details><summary>Radix Sort</summary>
./sort/radixsort.go#L3-L74
</details>

<details><summary>Counting Sort</summary>
./sort/countingsort.go#L3-L59
</details>

<details><summary>Cube Sort</summary>
./sort/cubesort.go#L3-L46
</details>