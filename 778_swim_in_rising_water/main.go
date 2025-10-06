package main

import "container/heap"

type Cell struct {
	r, c, t int
}

type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].t < pq[j].t
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

func swimInWater(grid [][]int) int {
	n := len(grid)
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	isWithinBound := func(r, c int) bool {
		return r >= 0 && r < n && c >= 0 && c < n
	}

	pq := &PriorityQueue{{0, 0, grid[0][0]}}
	heap.Init(pq)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(*Cell)
		if cell.r == n-1 && cell.c == n-1 {
			return cell.t
		}

		if visited[cell.r][cell.c] {
			continue
		}

		visited[cell.r][cell.c] = true
		for _, dir := range directions {
			nr, nc := cell.r+dir[0], cell.c+dir[1]
			if !isWithinBound(nr, nc) {
				continue
			}

			heap.Push(pq, &Cell{
				r: nr,
				c: nc,
				t: max(cell.t, grid[nr][nc]),
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
