package main

import "container/heap"

type MaxHeap [][]int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][2] > h[j][2] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	safenessMap := make(map[int]map[int]int) // row->col->dist

	// find all the thieves
	thieves := [][]int{}
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == 1 {
				thieves = append(thieves, []int{row, col, 0})
				if _, ok := safenessMap[row]; !ok {
					safenessMap[row] = make(map[int]int)
				}
				safenessMap[row][col] = 0
			}
		}
	}

	// multi sources BFS
	queue := thieves
	for len(queue) > 0 {
		newQueue := [][]int{}
		for _, cell := range queue {
			row, col, dist := cell[0], cell[1], cell[2]
			neighbours := [][]int{}

			// top
			if row > 0 {
				neighbours = append(neighbours, []int{row - 1, col, dist + 1})
			}

			// bottom
			if row < n-1 {
				neighbours = append(neighbours, []int{row + 1, col, dist + 1})
			}

			// left
			if col > 0 {
				neighbours = append(neighbours, []int{row, col - 1, dist + 1})
			}

			// right
			if col < n-1 {
				neighbours = append(neighbours, []int{row, col + 1, dist + 1})
			}

			// calculate safeness factor
			for _, neighbour := range neighbours {
				nRow, nCol, nDist := neighbour[0], neighbour[1], neighbour[2]
				if _, ok := safenessMap[nRow]; !ok {
					safenessMap[nRow] = make(map[int]int)
				}

				if _, ok := safenessMap[nRow][nCol]; !ok {
					safenessMap[nRow][nCol] = nDist
					newQueue = append(newQueue, neighbour)
				}
			}
		}
		queue = newQueue
	}

	// initialize max heap and visited hashmap
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, []int{0, 0, safenessMap[0][0]})
	visited := make(map[int]map[int]bool)
	visited[0] = make(map[int]bool)
	visited[0][0] = true

	for maxHeap.Len() > 0 {
		cell := heap.Pop(maxHeap).([]int)
		row, col, dist := cell[0], cell[1], cell[2]

		// return if we reach the destination
		if row == n-1 && col == n-1 {
			return dist
		}

		neighbours := [][]int{}

		// top
		if row > 0 {
			neighbours = append(neighbours, []int{row - 1, col})
		}

		// bottom
		if row < n-1 {
			neighbours = append(neighbours, []int{row + 1, col})
		}

		// left
		if col > 0 {
			neighbours = append(neighbours, []int{row, col - 1})
		}

		// right
		if col < n-1 {
			neighbours = append(neighbours, []int{row, col + 1})
		}

		// visit the neighbours
		for _, neighbour := range neighbours {
			nRow, nCol := neighbour[0], neighbour[1]

			// check if visited
			if _, ok := visited[nRow]; !ok {
				visited[nRow] = make(map[int]bool)
			}
			if visited, ok := visited[nRow][nCol]; ok && visited {
				continue
			}

			visited[nRow][nCol] = true
			minDist := min(dist, safenessMap[nRow][nCol])
			heap.Push(maxHeap, []int{nRow, nCol, minDist})
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
