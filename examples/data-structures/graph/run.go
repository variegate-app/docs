package main

import "fmt"

type Graph struct {
	adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adj: map[int][]int{}}
}

// N/A: граф обычно не адресуется как массив.
func (g *Graph) Access(_ int) (int, bool) {
	return 0, false
}

// O(V + E): BFS/DFS в общем случае проходит вершины и ребра.
func (g *Graph) Search(target int) bool {
	visited := map[int]bool{}
	for start := range g.adj {
		if visited[start] {
			continue
		}
		q := []int{start}
		visited[start] = true
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			if v == target {
				return true
			}
			for _, to := range g.adj[v] {
				if !visited[to] {
					visited[to] = true
					q = append(q, to)
				}
			}
		}
	}
	return false
}

// O(1): добавление ребра в список смежности (амортиз.).
func (g *Graph) Insert(from, to int) {
	g.adj[from] = append(g.adj[from], to)
	if _, ok := g.adj[to]; !ok {
		g.adj[to] = nil
	}
}

// O(V + E): удаление вершины и входящих ребер.
func (g *Graph) Delete(vertex int) {
	delete(g.adj, vertex)
	for from := range g.adj {
		filtered := g.adj[from][:0]
		for _, to := range g.adj[from] {
			if to != vertex {
				filtered = append(filtered, to)
			}
		}
		g.adj[from] = filtered
	}
}

func main() {
	g := NewGraph()
	g.Insert(1, 2)
	g.Insert(2, 3)
	g.Insert(3, 4)
	fmt.Println("Search(4):", g.Search(4))
	g.Delete(3)
	fmt.Println("Search(4) after Delete(3):", g.Search(4))
}
