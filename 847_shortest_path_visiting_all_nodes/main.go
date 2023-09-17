package main

func shortestPathLength(graph [][]int) int {

	// number of nodes
	n := len(graph)

	// queue of nodes to visit
	queue := [][2]int{}

	// visited nodes
	visited := make([][]bool, n)

	// initialize visited
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, 1<<n)
		visited[i][1<<i] = true
		queue = append(queue, [2]int{i, 1 << i})
	}

	// BFS
	for path := 0; ; path++ {
		for size := len(queue); size > 0; size-- {
			node := queue[0]
			queue = queue[1:]

			// return path length if all nodes are visited
			if node[1] == 1<<n-1 {
				return path
			}

			// visit all neighbors
			for _, neighbor := range graph[node[0]] {
				// visit unvisited neighbors
				if !visited[neighbor][node[1]|1<<neighbor] {
					visited[neighbor][node[1]|1<<neighbor] = true
					queue = append(queue, [2]int{neighbor, node[1] | 1<<neighbor})
				}
			}
		}
	}
}
