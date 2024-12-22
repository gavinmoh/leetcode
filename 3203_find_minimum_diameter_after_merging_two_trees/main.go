package main

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	n := len(edges1) + 1
	m := len(edges2) + 1

	buildAdjList := func(size int, edges [][]int) [][]int {
		adjList := make([][]int, size)
		for i := 0; i < size; i++ {
			adjList[i] = []int{}
		}
		for _, edge := range edges {
			adjList[edge[0]] = append(adjList[edge[0]], edge[1])
			adjList[edge[1]] = append(adjList[edge[1]], edge[0])
		}
		return adjList
	}

	adjList1 := buildAdjList(n, edges1)
	adjList2 := buildAdjList(m, edges2)

	findFarthestNode := func(adjList [][]int, sourceNode int) (int, int) {
		queue := []int{sourceNode}
		visited := map[int]bool{}
		visited[sourceNode] = true

		maximumDistance := 0
		farthestNode := sourceNode

		for len(queue) > 0 {
			newQueue := []int{}
			for _, currentNode := range queue {
				farthestNode = currentNode

				for _, neighbour := range adjList[currentNode] {
					if _, ok := visited[neighbour]; !ok {
						visited[neighbour] = true
						newQueue = append(newQueue, neighbour)
					}
				}
			}

			if len(newQueue) > 0 {
				maximumDistance++
			}
			queue = newQueue
		}

		return farthestNode, maximumDistance
	}

	findDiameter := func(n int, adjList [][]int) int {
		farthestNode, _ := findFarthestNode(adjList, 0)
		_, diameter := findFarthestNode(adjList, farthestNode)

		return diameter
	}

	diameter1 := findDiameter(n, adjList1)
	diameter2 := findDiameter(m, adjList2)

	combinedDiameter := ceilDivide(diameter1, 2) + ceilDivide(diameter2, 2) + 1

	return max(diameter1, max(diameter2, combinedDiameter))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func ceilDivide(x int, d int) int {
	if x%d == 0 {
		return x / d
	}

	return (x / d) + 1
}
