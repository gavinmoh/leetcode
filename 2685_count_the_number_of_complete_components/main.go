package main

func countCompleteComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = -1
		size[i] = 1
	}

	var find func(i int) int
	find = func(i int) int {
		if parent[i] == -1 {
			return i
		}

		parent[i] = find(parent[i])

		return parent[i]
	}

	union := func(i, j int) {
		left, right := find(i), find(j)
		if left == right {
			return
		}

		if size[left] > size[right] {
			parent[right] = left
			size[left] += size[right]
		} else {
			parent[left] = right
			size[right] += size[left]
		}
	}

	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	edgeCount := make([]int, n)
	for _, edge := range edges {
		root := find(edge[0])
		edgeCount[root] = edgeCount[root] + 1
	}

	completeCount := 0
	for i := 0; i < n; i++ {
		if find(i) == i {
			nodeCount := size[i]
			expectedEdges := nodeCount * (nodeCount - 1) / 2
			if edgeCount[i] == expectedEdges {
				completeCount++
			}
		}
	}

	return completeCount
}
