package main

import "sort"

func getAncestors(n int, edges [][]int) [][]int {
	nodes := make([][]int, n)
	for _, edge := range edges {
		parent, child := edge[0], edge[1]
		nodes[child] = append(nodes[child], parent)
	}

	var dfs func(node int, visited map[int]bool) []int
	dfs = func(node int, visited map[int]bool) []int {
		parents := nodes[node]
		ancestors := []int{}
		for _, parent := range parents {
			if _, ok := visited[parent]; !ok {
				ancestors = append(ancestors, parent)
				visited[parent] = true
				ancestors = append(ancestors, dfs(parent, visited)...)
			}
		}
		return ancestors
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		ancestors := dfs(i, make(map[int]bool))
		if len(ancestors) > 1 {
			sort.Ints(ancestors)
		}
		result[i] = ancestors
	}

	return result
}
