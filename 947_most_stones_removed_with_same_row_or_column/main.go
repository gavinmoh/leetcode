package main

func removeStones(stones [][]int) int {
	n := len(stones)
	count := n
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
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)

		if rootX == rootY {
			return
		}

		count--
		parent[rootX] = find(rootY)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				union(i, j)
			}
		}
	}

	return n - count
}
