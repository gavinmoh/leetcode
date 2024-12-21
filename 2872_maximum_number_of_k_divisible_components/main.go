package main

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adjacencyList := map[int][]int{}
	for i := 0; i < n; i++ {
		adjacencyList[i] = []int{}
	}

	for _, edge := range edges {
		left, right := edge[0], edge[1]
		adjacencyList[left] = append(adjacencyList[left], right)
		adjacencyList[right] = append(adjacencyList[right], left)
	}

	componentCount := 0

	var dfs func(current int, parent int) int
	dfs = func(current int, parent int) int {
		sum := 0
		for _, neighbour := range adjacencyList[current] {
			if neighbour != parent {
				sum += dfs(neighbour, current)
			}
			sum %= k
		}

		sum += values[current]
		sum %= k

		if sum == 0 {
			componentCount++
		}

		return sum
	}
	dfs(0, -1)

	return componentCount
}
