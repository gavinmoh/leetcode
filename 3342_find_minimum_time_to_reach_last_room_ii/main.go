package main

import (
	"container/heap"
	"math"
)

type Item struct {
	Row            int
	Col            int
	Cost           int
	PrevMovingTime int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
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

func minTimeToReach(moveTime [][]int) int {
	rows, cols := len(moveTime), len(moveTime[0])
	visited := make([][]bool, rows)
	costs := make([][]int, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
		costs[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			costs[i][j] = math.MaxInt
		}
	}

	queue := &PriorityQueue{}
	heap.Init(queue)
	heap.Push(queue, &Item{
		Row:            0,
		Col:            0,
		Cost:           0,
		PrevMovingTime: 2,
	})

	for queue.Len() > 0 {
		currentCell := heap.Pop(queue).(*Item)
		directions := [][]int{
			{0, -1}, // left
			{0, 1},  // right
			{-1, 0}, // top
			{1, 0},  // bottom
		}

		for _, direction := range directions {
			newRow, newCol := currentCell.Row+direction[0], currentCell.Col+direction[1]
			if newRow < 0 || newRow == rows || newCol < 0 || newCol == cols {
				continue
			}

			if visited[newRow][newCol] {
				continue
			}

			visited[newRow][newCol] = true

			movingTime := 1
			if currentCell.PrevMovingTime == 1 {
				movingTime += 1
			}

			newCost := max(currentCell.Cost, moveTime[newRow][newCol]) + movingTime
			if newCost < costs[newRow][newCol] {
				costs[newRow][newCol] = newCost
				heap.Push(queue, &Item{
					Row:            newRow,
					Col:            newCol,
					Cost:           newCost,
					PrevMovingTime: movingTime,
				})
			}
		}
	}

	return costs[rows-1][cols-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
