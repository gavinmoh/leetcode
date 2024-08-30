package main

const INF int = 2e9

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	dijkstra := func(edges [][]int) int {
		graph := make([][]int, n)
		distance := make([]int, n)
		visited := make([]bool, n)

		for i := 0; i < n; i++ {
			graph[i] = make([]int, n)
			for j := 0; j < n; j++ {
				graph[i][j] = INF
			}
			distance[i] = INF
		}

		distance[source] = 0
		for _, edge := range edges {
			a, b, weight := edge[0], edge[1], edge[2]
			if weight == -1 {
				continue
			}
			graph[a][b] = weight
			graph[b][a] = weight
		}

		for i := 0; i < n; i++ {
			k := -1
			for j := 0; j < n; j++ {
				if !visited[j] && (k == -1 || distance[j] < distance[k]) {
					k = j
				}
			}
			visited[k] = true
			for j := 0; j < n; j++ {
				distance[j] = min(distance[j], distance[k]+graph[k][j])
			}
		}
		return distance[destination]
	}
	distance := dijkstra(edges)
	if distance < target {
		return [][]int{}
	}

	ok := (distance == target)
	for _, edge := range edges {
		if edge[2] > 0 {
			continue
		}
		if ok {
			edge[2] = INF
			continue
		}
		edge[2] = 1
		distance := dijkstra(edges)
		if distance <= target {
			ok = true
			edge[2] += target - distance
		}
	}
	if ok {
		return edges
	}
	return [][]int{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
