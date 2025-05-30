package main

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	distanceFromNode1 := make([]int, n)
	distanceFromNode2 := make([]int, n)
	for i := 0; i < n; i++ {
		distanceFromNode1[i] = -1
		distanceFromNode2[i] = -1
	}

	bfs := func(node int, distances []int) {
		queue := []int{node}
		visited := make([]bool, n)
		dist := 0
		for len(queue) > 0 {
			newQueue := []int{}
			for _, node := range queue {
				if visited[node] {
					continue
				}
				visited[node] = true
				distances[node] = dist

				if edges[node] != -1 {
					newQueue = append(newQueue, edges[node])
				}
			}

			dist++
			queue = newQueue
		}
	}

	bfs(node1, distanceFromNode1)
	bfs(node2, distanceFromNode2)

	maxDistances := make([]int, n)
	for i := 0; i < n; i++ {
		if distanceFromNode1[i] == -1 || distanceFromNode2[i] == -1 {
			maxDistances[i] = -1
		} else {
			maxDistances[i] = max(distanceFromNode1[i], distanceFromNode2[i])
		}
	}

	answer := -1
	currentMax := -1

	for node, maxDistance := range maxDistances {
		if maxDistance == -1 {
			continue
		}

		if currentMax == -1 || maxDistance < currentMax {
			answer = node
			currentMax = maxDistance
		}
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
