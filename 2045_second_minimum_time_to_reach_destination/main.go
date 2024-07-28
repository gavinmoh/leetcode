package main

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make(map[int][]int)
	for _, edge := range edges {
		v1, v2 := edge[0], edge[1]
		if _, ok := adj[v1]; !ok {
			adj[v1] = []int{}
		}
		adj[v1] = append(adj[v1], v2)
		if _, ok := adj[v2]; !ok {
			adj[v2] = []int{}
		}
		adj[v2] = append(adj[v2], v1)
	}

	currentTime := 0
	queue := []int{1}
	visitedTimes := make(map[int][]int)
	result := -1
	for len(queue) > 0 {
		newQueue := []int{}
		for _, node := range queue {
			if node == n {
				if result != -1 {
					return currentTime
				}
				result = currentTime
			}

			for _, neighbour := range adj[node] {
				if _, ok := visitedTimes[neighbour]; !ok {
					visitedTimes[neighbour] = []int{}
				}

				if (len(visitedTimes[neighbour]) == 0) || (len(visitedTimes[neighbour]) == 1 && visitedTimes[neighbour][0] != currentTime) {
					newQueue = append(newQueue, neighbour)
					visitedTimes[neighbour] = append(visitedTimes[neighbour], currentTime)
				}
			}

		}

		if (currentTime/change)%2 != 0 {
			currentTime += (change - (currentTime % change))
		}
		currentTime += time
		queue = newQueue
	}

	return -1
}
