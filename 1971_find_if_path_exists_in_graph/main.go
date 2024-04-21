package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	// initialize parent
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	// find function
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			return find(parent[x])
		}
		return x
	}

	// union function
	var union func(x, y int)
	union = func(x, y int) {
		parent[find(y)] = find(x)
	}

	// union all the edges
	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	return find(source) == find(destination)
}
