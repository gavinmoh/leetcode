package main

import "container/heap"

type Edge struct {
	Row  int
	Col  int
	Cost int
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Edge)
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

func minimumObstacles(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	cells := &PriorityQueue{}
	heap.Init(cells)
	heap.Push(cells, &Edge{
		Row:  0,
		Col:  0,
		Cost: grid[0][0],
	})

	visited := map[int]map[int]bool{}

	for cells.Len() > 0 {
		cell := heap.Pop(cells).(*Edge)
		row, col, cost := cell.Row, cell.Col, cell.Cost
		if row == rows-1 && col == cols-1 {
			return cost
		}

		if _, ok := visited[row]; !ok {
			visited[row] = map[int]bool{}
		}

		if _, ok := visited[row][col]; ok {
			continue
		} else {
			visited[row][col] = true
		}

		// up
		if row > 0 {
			heap.Push(cells, &Edge{
				Row:  row - 1,
				Col:  col,
				Cost: cost + grid[row-1][col],
			})
		}
		// down
		if row < rows-1 {
			heap.Push(cells, &Edge{
				Row:  row + 1,
				Col:  col,
				Cost: cost + grid[row+1][col],
			})
		}
		// left
		if col > 0 {
			heap.Push(cells, &Edge{
				Row:  row,
				Col:  col - 1,
				Cost: cost + grid[row][col-1],
			})
		}
		// right
		if col < cols-1 {
			heap.Push(cells, &Edge{
				Row:  row,
				Col:  col + 1,
				Cost: cost + grid[row][col+1],
			})
		}
	}

	return -1
}
