package main

import "container/heap"

type Cell struct {
	Row    int
	Col    int
	Height int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Height < pq[j].Height
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	cell := x.(*Cell)
	*pq = append(*pq, cell)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	cell := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return cell
}

func trapRainWater(heightMap [][]int) int {
	dRow := []int{0, 0, -1, 1}
	dCol := []int{-1, 1, 0, 0}
	rows, cols := len(heightMap), len(heightMap[0])

	visited := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		visited[row] = make([]bool, cols)
	}

	boundary := &PriorityQueue{}
	heap.Init(boundary)

	// push boundary cells
	for row := 0; row < rows; row++ {
		// first or last row
		if row == 0 || row == rows-1 {
			for col := 0; col < cols; col++ {
				height := heightMap[row][col]
				heap.Push(boundary, &Cell{
					Row:    row,
					Col:    col,
					Height: height,
				})
				visited[row][col] = true
			}
		} else {
			// first column
			heap.Push(boundary, &Cell{
				Row:    row,
				Col:    0,
				Height: heightMap[row][0],
			})
			visited[row][0] = true

			// last column
			heap.Push(boundary, &Cell{
				Row:    row,
				Col:    cols - 1,
				Height: heightMap[row][cols-1],
			})
			visited[row][cols-1] = true
		}
	}

	isValidCell := func(row, col int) bool {
		return row >= 0 && row < rows && col >= 0 && col < cols
	}

	totalWaterVolume := 0

	for boundary.Len() > 0 {
		cell := heap.Pop(boundary).(*Cell)
		minBoundaryHeight := cell.Height
		currentRow, currentCol := cell.Row, cell.Col

		for direction := 0; direction < 4; direction++ {
			neighbourRow := currentRow + dRow[direction]
			neighbourCol := currentCol + dCol[direction]

			if isValidCell(neighbourRow, neighbourCol) && !visited[neighbourRow][neighbourCol] {
				neighbourHeight := heightMap[neighbourRow][neighbourCol]
				if neighbourHeight < minBoundaryHeight {
					totalWaterVolume += minBoundaryHeight - neighbourHeight
				}
				heap.Push(boundary, &Cell{
					Row:    neighbourRow,
					Col:    neighbourCol,
					Height: max(neighbourHeight, minBoundaryHeight),
				})
				visited[neighbourRow][neighbourCol] = true
			}
		}
	}

	return totalWaterVolume
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
