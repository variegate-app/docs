## [<<< ---](../README.md)

## Основные типы алгоритмов:

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

### Сортировка

| Алгоритм | Time Best| Time Middle | Time Worst | Space Worst|
|---|---|---|---|---|
| [Quick sort](./sort/quick-sort.md) |🔴 O(n log n)|🔴 O(n log n)|🔴 O(n)|🟢 O(n*log(n))|
| [Merge sort](./sort/merge-sort.md) |🔴 O(n log n)|🔴 O(n log n)|🔴 O(n log n)| 🔵 O(n)|
| [Tim sort](./sort/tim-sort.md) |🔵 O(n)|🔴 O(n log n)|🔴 O(n log n)|🔵 O(1)|
| [Heap sort](./sort/heap-sort.md) |🔴 O(n log n)|🔴 O(n log n)|🔴 O(n log n)|🟢 O(1)|
| [Bubble sort](./sort/bubble-sort.md) |🔵 O(n)|🔴 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🟢 O(1)|
| [Insertion Sort](./sort/insertion-sort.md) |🔵 O(n)|🔴 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🟢 O(1)|
| [Selection Sort](./sort/selection-sort.md) |🟢 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🔴 O(n<sup>2</sup>)|🟢 O(1)|
| [Tree Sort](./sort/tree-sort.md) |🔴 O(n log n)|🔴 O(n log(n) )|🔴 O(n<sup>2</sup>) * T|🔵 O(n)|
| [Shell Sort](./sort/shell-sort.md) |🔴 O(n log n)|🔴 O(n (log n)<sup>2</sup>)|🔴 O(n (log n)<sup>2</sup>)|🟢 O(1)|
| [Bucket Sort](./sort/bucket-sort.md) |🟢 O(n + K)|🟢 O(n + K)|🔴 O(n<sup>2</sup>)|🔵 O(n)|
| [Radix Sort](./sort/radix-sort.md) |🟢 O( nK )|🟢 O( nK )|🟢 O( nK )|🔵 O( n + K )|
| [Counting Sort](./sort/counting-sort.md)  |🟢 O(n + K)|🟢 O(n + K)|🟢 O(n + K)| 🔵 O(K)|
| [Cube Sort](./sort/cube-sort.md)  |🔵 O(n)|🔴 O(n log n)|🔴 O(n log n)|🔵 O(n)|

### Структуры данных

| Структура данных | Access | Search | Insertion | Deletion | Space complexity |
|---|---|---|---|---|---|
| [Array](./data-structures/array.md) | 🟢 O(1) | 🔴 O(n) | 🔴 O(n) | 🔴 O(n) | 🔵 O(n) |
| [Queue](./data-structures/queue.md) | 🟢 O(1) | 🔴 O(n) | 🟢 O(1) | 🟢 O(1) | 🔵 O(n) |
| [Hash Table](./data-structures/hash-table.md) | N/A | 🟢 O(1)\* | 🟢 O(1)\* | 🟢 O(1)\* | 🔵 O(n) |
| [Graph](./data-structures/graph.md) | N/A | 🔴 O(V) | 🟢 O(1) | 🔴 O(V + E) | 🔵 O(V + E) |
| [Stack](./data-structures/stack.md) | 🟢 O(1) | 🔴 O(n) | 🟢 O(1) | 🟢 O(1) | 🔵 O(n) |
| [Heap](./data-structures/heap.md) | 🔵 O(1) | 🔴 O(n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |
| [Singly Linked List](./data-structures/singly-linked-list.md) | 🔴 O(n) | 🔴 O(n) | 🟢 O(1) | 🟢 O(1) | 🔵 O(n) |
| [Doubly Linked List](./data-structures/doubly-linked-list.md) | 🔴 O(n) | 🔴 O(n) | 🟢 O(1) | 🟢 O(1) | 🔵 O(n) |
| [Skip List](./data-structures/skip-list.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(n*log(n))|
| [Tree](./data-structures/tree.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |
| [Cartesian Tree](./data-structures/cartesian-tree.md) | 🔴 O(n) | 🔴 O(n) | 🔴 O(n) | 🔴 O(n) | 🔵 O(n) |
| [B-Tree](./data-structures/b-tree.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |
| [Binary Tree](./data-structures/binary-tree.md) | 🔴 O(n) | 🔴 O(n) | 🔴 O(n) | 🔴 O(n) | 🔵 O(n) |
| [Binary Search Tree](./data-structures/binary-search-tree.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |
| [Red-Black Tree](./data-structures/red-black-tree.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |
| [Splay Tree](./data-structures/splay-tree.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |
| [AVL Tree](./data-structures/avl-tree.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |
| [KD Tree](./data-structures/kd-tree.md) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔴 O(log n) | 🔵 O(n) |