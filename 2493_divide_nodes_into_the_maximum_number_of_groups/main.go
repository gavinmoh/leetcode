package main

func magnificentSets(n int, edges [][]int) int {
	adjList := make([][]int, n+1)
	parent, depth := make([]int, n+1), make([]int, n+1)
	for i := 0; i < n+1; i++ {
		adjList[i] = []int{}
		parent[i] = -1
	}

	// union find algorithm
	var find func(node int) int
	find = func(node int) int {
		for parent[node] != -1 {
			node = parent[node]
		}
		return node
	}

	var union func(node1, node2 int)
	union = func(node1, node2 int) {
		node1 = find(node1)
		node2 = find(node2)

		if node1 == node2 {
			return
		}

		if depth[node1] < depth[node2] {
			node1, node2 = node2, node1
		}
		parent[node2] = node1

		if depth[node1] == depth[node2] {
			depth[node1]++
		}
	}

	var getNumOfGroups func(node int) int
	getNumOfGroups = func(node int) int {
		nodesQueue := []int{}
		layerSeen := make([]int, n+1)
		for i := 0; i < n+1; i++ {
			layerSeen[i] = -1
		}
		nodesQueue = append(nodesQueue, node)
		layerSeen[node] = 0
		deepestLayer := 0

		// bfs
		for len(nodesQueue) > 0 {
			newNodesQueue := []int{}
			for _, currentNode := range nodesQueue {
				for _, neighbour := range adjList[currentNode] {
					if layerSeen[neighbour] == -1 {
						layerSeen[neighbour] = deepestLayer + 1
						newNodesQueue = append(newNodesQueue, neighbour)
					} else {
						if layerSeen[neighbour] == deepestLayer {
							return -1
						}
					}
				}
			}
			deepestLayer++
			nodesQueue = newNodesQueue
		}

		return deepestLayer
	}

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]
		adjList[node2] = append(adjList[node2], node1)
		adjList[node1] = append(adjList[node1], node2)
		union(node1, node2)
	}

	numOfGroupsForComponent := map[int]int{}
	for node := 1; node < n+1; node++ {
		numOfGroups := getNumOfGroups(node)
		if numOfGroups == -1 {
			return -1
		}
		rootNode := find(node)
		if _, ok := numOfGroupsForComponent[rootNode]; !ok {
			numOfGroupsForComponent[rootNode] = numOfGroups
		} else {
			numOfGroupsForComponent[rootNode] = max(
				numOfGroupsForComponent[rootNode],
				numOfGroups,
			)
		}
	}

	totalNumOfGroups := 0
	for _, numOfGroups := range numOfGroupsForComponent {
		totalNumOfGroups += numOfGroups
	}

	return totalNumOfGroups
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
