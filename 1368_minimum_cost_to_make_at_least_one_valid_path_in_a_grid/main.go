package main

import (
	"container/heap"
	"math"
)

type Cell struct {
	Row  int
	Col  int
	Cost int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
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

func minCost(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	queue := &PriorityQueue{}
	heap.Init(queue)
	heap.Push(queue, &Cell{
		Row:  0,
		Col:  0,
		Cost: 0,
	})

	minCost := make([][]int, rows)
	for i := 0; i < rows; i++ {
		minCost[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			minCost[i][j] = math.MaxInt
		}
	}
	minCost[0][0] = 0

	for queue.Len() > 0 {
		cell := heap.Pop(queue).(*Cell)
		row, col, cost := cell.Row, cell.Col, cell.Cost
		sign := grid[row][col]

		if minCost[row][col] != cost {
			continue
		}

		// right
		if col < cols-1 {
			newCost := cost + calculateCost(sign, 1)
			if minCost[row][col+1] > newCost {
				minCost[row][col+1] = newCost
				heap.Push(queue, &Cell{
					Row:  row,
					Col:  col + 1,
					Cost: newCost,
				})
			}
		}

		// left
		if col > 0 {
			newCost := cost + calculateCost(sign, 2)
			if minCost[row][col-1] > newCost {
				minCost[row][col-1] = newCost
				heap.Push(queue, &Cell{
					Row:  row,
					Col:  col - 1,
					Cost: newCost,
				})
			}
		}

		// lower
		if row < rows-1 {
			newCost := cost + calculateCost(sign, 3)
			if minCost[row+1][col] > newCost {
				minCost[row+1][col] = newCost
				heap.Push(queue, &Cell{
					Row:  row + 1,
					Col:  col,
					Cost: newCost,
				})
			}
		}

		// upper
		if row > 0 {
			newCost := cost + calculateCost(sign, 4)
			if minCost[row-1][col] > newCost {
				minCost[row-1][col] = newCost
				heap.Push(queue, &Cell{
					Row:  row - 1,
					Col:  col,
					Cost: newCost,
				})
			}
		}
	}

	return minCost[rows-1][cols-1]
}

func calculateCost(sign, direction int) int {
	if sign == direction {
		return 0
	}
	return 1
}
