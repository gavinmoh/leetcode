package main

import "container/heap"

type Query struct {
	MaxHeight int
	Index     int
}

type PriorityQueue []*Query

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].MaxHeight < pq[j].MaxHeight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	query := x.(*Query)
	*pq = append(*pq, query)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	query := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return query
}

func (pq *PriorityQueue) Top() any {
	arr := *pq
	return arr[0]
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	storeQueries := map[int][]*Query{}
	// create and initialize ans slice to -1
	ans := make([]int, len(queries))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}

	for i, query := range queries {
		x, y := query[0], query[1]
		if x < y && heights[x] < heights[y] {
			ans[i] = y
		} else if x > y && heights[x] > heights[y] {
			ans[i] = x
		} else if x == y {
			ans[i] = x
		} else {
			if _, ok := storeQueries[max(x, y)]; !ok {
				storeQueries[max(x, y)] = []*Query{}
			}

			storeQueries[max(x, y)] = append(storeQueries[max(x, y)], &Query{
				Index:     i,
				MaxHeight: max(heights[x], heights[y]),
			})
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for i, height := range heights {
		for pq.Len() > 0 && pq.Top().(*Query).MaxHeight < height {
			query := heap.Pop(pq).(*Query)
			ans[query.Index] = i
		}

		for _, query := range storeQueries[i] {
			heap.Push(pq, query)
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
