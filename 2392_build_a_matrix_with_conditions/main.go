package main

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowOrder := topoSort(rowConditions, k)
	colOrder := topoSort(colConditions, k)

	if len(rowOrder) == 0 || len(colOrder) == 0 {
		return [][]int{}
	}

	// reverse and convert to hashmap
	rowMap := make(map[int]int)
	for i, num := range rowOrder {
		rowMap[num] = k - i - 1
	}
	colMap := make(map[int]int)
	for i, num := range colOrder {
		colMap[num] = k - i - 1
	}

	// initialize the matrix
	matrix := make([][]int, k)
	for i := 0; i < k; i++ {
		matrix[i] = make([]int, k)
	}

	for i := 1; i <= k; i++ {
		row, col := rowMap[i], colMap[i]
		matrix[row][col] = i
	}

	return matrix
}

func topoSort(graphs [][]int, k int) []int {
	visited := make(map[int]bool)
	adjacencyList := make(map[int][]int)
	result := []int{}

	for _, graph := range graphs {
		parent, child := graph[0], graph[1]
		if _, ok := adjacencyList[parent]; !ok {
			adjacencyList[parent] = []int{}
		}
		adjacencyList[parent] = append(adjacencyList[parent], child)
	}

	var dfs func(node int, path map[int]bool) bool
	dfs = func(node int, path map[int]bool) bool {
		if path[node] {
			return false
		}
		if visited[node] {
			return true
		}
		visited[node] = true

		if _, ok := adjacencyList[node]; !ok {
			result = append(result, node)
			return true
		}

		for _, child := range adjacencyList[node] {
			newPath := make(map[int]bool)
			for k, v := range path {
				newPath[k] = v
			}
			newPath[node] = true
			if !dfs(child, newPath) {
				return false
			}
		}

		result = append(result, node)
		return true
	}

	for i := 1; i <= k; i++ {
		path := make(map[int]bool)
		if !dfs(i, path) {
			return []int{}
		}
	}

	return result
}
