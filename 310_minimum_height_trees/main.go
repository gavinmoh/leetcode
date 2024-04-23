package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	nodes := make(map[int][]int)
	edgeCount := make([]int, n)
	for _, edge := range edges {
		left, right := edge[0], edge[1]
		if _, ok := nodes[left]; !ok {
			nodes[left] = []int{right}
		} else {
			nodes[left] = append(nodes[left], right)
		}
		if _, ok := nodes[right]; !ok {
			nodes[right] = []int{left}
		} else {
			nodes[right] = append(nodes[right], left)
		}
		edgeCount[left]++
		edgeCount[right]++
	}

	queue := []int{}
	for node, count := range edgeCount {
		if count == 1 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 && n > 2 {
		newQueue := []int{}
		for _, node := range queue {
			for _, neighbour := range nodes[node] {
				edgeCount[neighbour]--
				if edgeCount[neighbour] == 1 {
					newQueue = append(newQueue, neighbour)
				}
			}
		}
		n -= len(queue)
		queue = newQueue
	}

	return queue
}
