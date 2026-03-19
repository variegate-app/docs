package main

import (
	"fmt"
)

// HamiltonianCycle finds an exact Hamiltonian cycle in the given graph.
// It is an "optimal" decision solver: if a Hamiltonian cycle exists,
// it returns one; otherwise it returns (nil, false).
//
// Complexity is exponential (Hamiltonian cycle is NP-complete),
// but this backtracking is correct for the small graphs typically used
// in documentation examples.
func HamiltonianCycle(adj [][]bool) ([]int, bool) {
	n := len(adj)
	if n == 0 {
		return nil, false
	}

	// Fix a start vertex to avoid equivalent rotations of the same cycle.
	start := 0

	path := make([]int, n)
	visited := make([]bool, n)
	path[0] = start
	visited[start] = true

	var dfs func(pos int) ([]int, bool)
	dfs = func(pos int) ([]int, bool) {
		// Build positions 1..n-1 (position pos is the next slot to fill).
		if pos == n {
			// Close the cycle: last -> start.
			if adj[path[n-1]][start] {
				return append([]int(nil), path...), true
			}
			return nil, false
		}

		// Try next vertices.
		for next := 0; next < n; next++ {
			if visited[next] {
				continue
			}
			if !adj[path[pos-1]][next] {
				continue
			}

			visited[next] = true
			path[pos] = next

			if res, ok := dfs(pos + 1); ok {
				return res, true
			}

			visited[next] = false
		}
		return nil, false
	}

	return dfs(1)
}

func main() {
	// Real-life example:
	// Delivery points (vertices) and allowed direct legs (edges).
	//
	// Vertices:
	// 0:A, 1:B, 2:C, 3:D, 4:E, 5:F
	//
	// Edges form a graph that contains a Hamiltonian cycle.

	labels := []string{"A", "B", "C", "D", "E", "F"}
	n := len(labels)

	adj := make([][]bool, n)
	for i := range adj {
		adj[i] = make([]bool, n)
	}

	// Undirected graph: set edges both directions.
	addEdge := func(u, v int) {
		adj[u][v] = true
		adj[v][u] = true
	}

	// Build edges (some extra legs + at least one Hamiltonian cycle).
	addEdge(0, 1) // A-B
	addEdge(1, 2) // B-C
	addEdge(2, 3) // C-D
	addEdge(3, 4) // D-E
	addEdge(4, 5) // E-F
	addEdge(5, 0) // F-A  => A-B-C-D-E-F-A is a Hamiltonian cycle

	// Extra allowed routes:
	addEdge(0, 2) // A-C
	addEdge(1, 3) // B-D
	addEdge(2, 4) // C-E
	addEdge(3, 5) // D-F

	cycle, ok := HamiltonianCycle(adj)
	if !ok {
		fmt.Println("No Hamiltonian cycle found.")
		return
	}

	fmt.Println("Hamiltonian cycle found:")
	for i, v := range cycle {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(labels[v])
	}
	fmt.Printf(" -> %s\n", labels[cycle[0]]) // close the cycle
}

