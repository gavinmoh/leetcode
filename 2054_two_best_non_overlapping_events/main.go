package main

import (
	"container/heap"
	"sort"
)

type Event struct {
	End   int
	Value int
}

type PriorityQueue []*Event

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].End < pq[j].End
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	event := x.(*Event)
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

func (pq *PriorityQueue) Top() any {
	arr := *pq
	return arr[0]
}

func maxTwoEvents(events [][]int) int {
	// sort events by start time
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	pq := &PriorityQueue{}
	heap.Init(pq)

	maxVal, maxSum := 0, 0
	for _, event := range events {
		start, end, value := event[0], event[1], event[2]
		for pq.Len() > 0 && pq.Top().(*Event).End < start {
			maxVal = max(maxVal, pq.Top().(*Event).Value)
			heap.Pop(pq)
		}

		maxSum = max(maxSum, maxVal+value)
		heap.Push(pq, &Event{
			End:   end,
			Value: value,
		})
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
