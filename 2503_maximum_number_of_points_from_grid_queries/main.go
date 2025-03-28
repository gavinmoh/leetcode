package main

import (
	"container/heap"
	"sort"
)

type Cell struct {
	Row   int
	Col   int
	Value int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
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

func maxPoints(grid [][]int, queries []int) []int {
	rows, cols := len(grid), len(grid[0])

	sortedQuery := [][]int{}
	for i, query := range queries {
		sortedQuery = append(sortedQuery, []int{i, query})
	}
	sort.SliceStable(sortedQuery, func(i, j int) bool {
		return sortedQuery[i][1] < sortedQuery[j][1]
	})

	answers := make([]int, len(queries))
	visited := make([][]bool, rows)
	for row := range visited {
		visited[row] = make([]bool, cols)
	}

	points := 0
	queue := &PriorityQueue{}
	heap.Init(queue)
	heap.Push(queue, &Cell{Row: 0, Col: 0, Value: grid[0][0]})
	for _, pair := range sortedQuery {
		i, query := pair[0], pair[1]

		for queue.Len() > 0 {
			cell := heap.Pop(queue).(*Cell)

			if visited[cell.Row][cell.Col] {
				continue
			}

			if cell.Value >= query {
				heap.Push(queue, cell)
				break
			}

			visited[cell.Row][cell.Col] = true
			points++

			if cell.Row > 0 {
				heap.Push(queue, &Cell{
					Row:   cell.Row - 1,
					Col:   cell.Col,
					Value: grid[cell.Row-1][cell.Col],
				})
			}

			if cell.Col > 0 {
				heap.Push(queue, &Cell{
					Row:   cell.Row,
					Col:   cell.Col - 1,
					Value: grid[cell.Row][cell.Col-1],
				})
			}

			if cell.Row < rows-1 {
				heap.Push(queue, &Cell{
					Row:   cell.Row + 1,
					Col:   cell.Col,
					Value: grid[cell.Row+1][cell.Col],
				})
			}

			if cell.Col < cols-1 {
				heap.Push(queue, &Cell{
					Row:   cell.Row,
					Col:   cell.Col + 1,
					Value: grid[cell.Row][cell.Col+1],
				})
			}
		}

		answers[i] = points
	}

	return answers
}
