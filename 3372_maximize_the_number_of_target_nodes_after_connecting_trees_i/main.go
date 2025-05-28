package main

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	n, m := len(edges1)+1, len(edges2)+1
	adjList1, adjList2 := make([][]int, n), make([][]int, m)
	for i := 0; i < n; i++ {
		adjList1[i] = []int{}
	}
	for i := 0; i < m; i++ {
		adjList2[i] = []int{}
	}

	for _, edge := range edges1 {
		u, v := edge[0], edge[1]
		adjList1[u] = append(adjList1[u], v)
		adjList1[v] = append(adjList1[v], u)
	}

	for _, edge := range edges2 {
		u, v := edge[0], edge[1]
		adjList2[u] = append(adjList2[u], v)
		adjList2[v] = append(adjList2[v], u)
	}

	var dfs func(node int, parent int, step int, adjList [][]int) int
	dfs = func(node int, parent int, step int, adjList [][]int) int {
		if step < 0 {
			return 0
		}

		count := 1
		for _, neighbour := range adjList[node] {
			if neighbour == parent {
				continue
			}
			count += dfs(neighbour, node, step-1, adjList)
		}

		return count
	}

	count1, count2 := make([]int, n), make([]int, m)
	for node := 0; node < n; node++ {
		count1[node] = dfs(node, -1, k, adjList1)
	}
	for node := 0; node < m; node++ {
		count2[node] = dfs(node, -1, k-1, adjList2)
	}

	maxCount2 := 0
	for _, count := range count2 {
		if count > maxCount2 {
			maxCount2 = count
		}
	}

	answer := make([]int, n)
	for node := 0; node < n; node++ {
		answer[node] = count1[node] + maxCount2
	}

	return answer
}
