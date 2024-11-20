package main

import "container/heap"

type Cell struct {
	Time int
	Row  int
	Col  int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Time < pq[j].Time
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Cell)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return item
}

func minimumTime(grid [][]int) int {
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	rows, cols := len(grid), len(grid[0])
	directions := [][]int{
		{-1, 0}, // top
		{1, 0},  // bottom
		{0, -1}, // left
		{0, 1},  // right
	}
	visited := map[int]map[int]bool{}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Cell{
		Time: grid[0][0],
		Row:  0,
		Col:  0,
	})

	isValid := func(row, col int) bool {
		if _, ok := visited[row][col]; ok {
			return true
		}

		return row >= 0 && row < rows && col >= 0 && col < cols
	}

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(*Cell)
		row, col, time := cell.Row, cell.Col, cell.Time
		if row == rows-1 && col == cols-1 {
			return time
		}

		if _, ok := visited[row]; !ok {
			visited[row] = map[int]bool{}
		}
		// skip if visited
		if _, ok := visited[row][col]; ok {
			continue
		} else {
			visited[row][col] = true
		}

		for _, direction := range directions {
			dy, dx := direction[0], direction[1]
			nextRow, nextCol := row+dy, col+dx
			if !isValid(nextRow, nextCol) {
				continue
			}

			waitTime := 0
			if (grid[nextRow][nextCol]-time)%2 == 0 {
				waitTime = 1
			}
			nextTime := max(grid[nextRow][nextCol]+waitTime, time+1)
			heap.Push(pq, &Cell{
				Time: nextTime,
				Row:  nextRow,
				Col:  nextCol,
			})
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
