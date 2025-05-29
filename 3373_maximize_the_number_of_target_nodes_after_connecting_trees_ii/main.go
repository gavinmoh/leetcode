package main

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	n, m := len(edges1)+1, len(edges2)+1
	colour1, colour2 := make([]int, n), make([]int, m)
	count1, count2 := countColour(edges1, colour1), countColour(edges2, colour2)

	answer := make([]int, n)
	for node := 0; node < n; node++ {
		answer[node] = count1[colour1[node]] + max(count2[0], count2[1])
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func buildAdjacencyList(edges [][]int) [][]int {
	n := len(edges) + 1
	adjList := make([][]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
		adjList[v] = append(adjList[v], u)
	}

	return adjList
}

func countColour(edges [][]int, colour []int) [2]int {
	n := len(edges) + 1
	adjList := buildAdjacencyList(edges)

	var dfs func(node, parent, distance int) int
	dfs = func(node, parent, distance int) int {
		count := 1 - (distance % 2)
		colour[node] = distance % 2
		for _, neighbour := range adjList[node] {
			if neighbour == parent {
				continue
			}
			count += dfs(neighbour, node, distance+1)
		}
		return count
	}

	evenCount := dfs(0, -1, 0)

	return [2]int{evenCount, n - evenCount}
}
