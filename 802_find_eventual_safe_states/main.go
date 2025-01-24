package main

func eventualSafeNodes(graph [][]int) []int {
	cache := map[int]bool{}

	var hasCycle func(node int, visited map[int]bool) bool
	hasCycle = func(node int, visited map[int]bool) bool {
		if _, ok := visited[node]; ok {
			return true
		}

		if cachedResult, ok := cache[node]; ok {
			return cachedResult
		}

		visited[node] = true
		result := false

		for _, nextNode := range graph[node] {
			if hasCycle(nextNode, visited) {
				result = true
				break
			}
		}

		delete(visited, node)
		cache[node] = result

		return result
	}

	cyclesNode := map[int]bool{}
	for node, nextNodes := range graph {
		visited := map[int]bool{node: true}
		for _, nextNode := range nextNodes {
			if hasCycle(nextNode, visited) {
				cyclesNode[node] = true
				break
			}
		}
	}

	result := []int{}
	for node := 0; node < len(graph); node++ {
		if _, ok := cyclesNode[node]; !ok {
			result = append(result, node)
		}
	}

	return result
}
