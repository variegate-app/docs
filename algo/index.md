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

### NP-полные задачи

Ниже — набор классических NP-полных задач (decision-версии: ответ `да/нет`), которые встречаются в учебных материалах и используются в редукциях.

#### Методы решения
Поскольку точные алгоритмы для NP-полных задач неэффективны на больших входных данных, исследователи разработали несколько стратегий для их решения: 
- Точные алгоритмы с оптимизацией — методы ветвей и границ, динамическое программирование.
- Приближённые алгоритмы — дают решение, близкое к оптимальному, с гарантированной погрешностью.
- Ограничение и параметризация — ограничивая область поиска или фиксируя определённые параметры, можно упростить NP-полную задачу до более управляемой.

Эти подходы не гарантируют нахождение оптимального решения, но в большинстве случаев они позволяют найти достаточно хорошие решения за разумное время. 

<details><summary>Булевы (SAT)</summary>

- [SAT (satisfiability)](./np-full/sat.md) — Выполнимость (SAT): есть ли назначение переменных, при котором формула истинна.
- [3SAT](./np-full/3sat.md) — 3-выполнимость (3SAT): есть ли назначение переменных для КНФ с clauses ровно по 3 литерала, при котором формула истинна.
- [MAX-SAT (decision)](./np-full/max-sat-decision.md) — Максимальная удовлетворимость: можно ли удовлетворить минимум `k` дизъюнктов.

</details>

<details><summary>Графовые (вершины/раскраска/разрезы)</summary>

- [VERTEX COVER](./np-full/vertex-cover.md) — Покрытие вершин: есть ли подмножество вершин размера `k`, инцидентное каждому ребру.
- [DOMINATING SET](./np-full/dominating-set.md) — Доминирующее множество: есть ли множество `k` вершин, такое что все вершины графа либо в нем, либо соседние с ним.
- [k-Colorability (graph coloring)](./np-full/k-colorability.md) — `k`-раскраска: можно ли раскрасить вершины в `k` цветов так, чтобы смежные вершины имели разные цвета.

</details>

<details><summary>Пути, циклы и туры</summary>

- [HAMILTONIAN PATH](./np-full/hamiltonian-path.md) — Гамильтонов путь: есть ли простой путь, проходящий по всем вершинам ровно один раз.
- [TRAVELING SALESMAN (decision)](./np-full/traveling-salesman-decision.md) — Продавец-путешественник (decision): существует ли гамильтонов цикл со стоимостью ≤ `B`.
- [k Disjoint Paths (decision)](./np-full/k-disjoint-paths.md) — k раздельных путей: можно ли связать пары вершин непересекающимися путями.

</details>

<details><summary>Связность и деревья</summary>

- [STEINER TREE (decision)](./np-full/steiner-tree-decision.md) — Дерево Штейнера (decision): есть ли подграф, соединяющий все терминалы, с весом ≤ `B`.
- [Steiner Forest (decision)](./np-full/steiner-forest-decision.md) — Лес Штейнера: можно ли соединить несколько терминальных множеств поддеревьями суммарной стоимости ≤ `B`.
- [Group Steiner Tree (decision)](./np-full/group-steiner-tree-decision.md) — Групповое дерево Штейнера: нужно подключить хотя бы одну вершину из каждой группы терминалов.

</details>

<details><summary>Покрытия, матчинг и упаковка</summary>

- [SET COVER](./np-full/set-cover.md) — Покрытие множествами: можно ли выбрать ≤ `k` подмножеств, объединение которых покрывает все элементы `U`.
- [BIN PACKING (decision)](./np-full/bin-packing-decision.md) — Бин-пэкинг (упаковка в бины, decision): разложить предметы по бинам емкости `C`, используя ≤ `k` бинов.
- [Hitting Set (decision)](./np-full/hitting-set-decision.md) — Хиттинг-множество: существует ли набор элементов, пересекающий каждое множество из семейства.

</details>

<details><summary>Числовые (суммы/разбиения/рюкзаки)</summary>

- [SUBSET SUM](./np-full/subset-sum.md) — Подмножество с суммой: существует ли подмножество элементов, дающее сумму ровно `T`.
- [0-1 KNAPSACK (decision)](./np-full/knapsack-01-decision.md) — Рюкзак 0/1 (decision): выбрать предметы так, чтобы суммарный вес ≤ `W` и ценность ≥ `V`.
- [PARTITION](./np-full/partition.md) — Разбиение: разбить набор на две части равной суммы.

</details>

<details><summary>Планирование и расписания</summary>

- [JOB SHOP SCHEDULING (decision)](./np-full/job-shop-scheduling-decision.md) — Планирование в job-shop (decision): существует ли расписание с makespan ≤ `B` при ограничениях по станкам и precedence.
- [Flow Shop Scheduling (decision)](./np-full/flow-shop-scheduling-decision.md) — Flow-shop (decision): есть ли расписание с makespan ≤ `B`.
- [RCPSP (Resource-Constrained Project Scheduling) (decision)](./np-full/rcpsp-decision.md) — RCPSP (decision): можно ли уложить все активности с учетом precedence и ограничений ресурсов в makespan ≤ `B`.

</details>
