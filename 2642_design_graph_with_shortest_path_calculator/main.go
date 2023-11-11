package main

type Graph struct {
	n         int
	distances [][]int // [from][to]cost
}

const infinity = 1 << 29

func Constructor(n int, edges [][]int) Graph {
	// initialize distances
	distances := make([][]int, n)
	for i := 0; i < n; i++ {
		distances[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				distances[i][j] = 0
			} else {
				distances[i][j] = infinity
			}
		}
	}

	// set the distances
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		cost := edge[2]
		distances[from][to] = cost
	}

	return Graph{n, distances}
}

func (this *Graph) AddEdge(edge []int) {
	from := edge[0]
	to := edge[1]
	cost := edge[2]
	this.distances[from][to] = cost
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	// Floyd-Warshall algorithm
	for k := 0; k < this.n; k++ {
		for i := 0; i < this.n; i++ {
			for j := 0; j < this.n; j++ {
				this.distances[i][j] = min(this.distances[i][j], this.distances[i][k]+this.distances[k][j])
			}
		}
	}

	if this.distances[node1][node2] == infinity {
		return -1
	}
	return this.distances[node1][node2]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
