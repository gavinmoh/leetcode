package main

import (
	"container/heap"
	"sort"
)

func maxEvents(events [][]int) int {
	n := len(events)
	maxDay := 0
	for _, event := range events {
		if event[1] > maxDay {
			maxDay = event[1]
		}
	}

	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	pq := &PriorityQueue{}
	heap.Init(pq)
	result := 0

	for i, day := 0, 1; day <= maxDay; day++ {
		for i < n && events[i][0] <= day {
			heap.Push(pq, events[i])
			i++
		}

		// pop expired event
		for pq.Len() > 0 && (*pq)[0][1] < day {
			heap.Pop(pq)
		}

		if pq.Len() > 0 {
			heap.Pop(pq)
			result++
		}
	}

	return result
}

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][1] < pq[j][1]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	event := x.([]int)
	*pq = append(*pq, event)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return event
}
