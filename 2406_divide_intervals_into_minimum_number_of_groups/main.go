package main

import (
	"container/heap"
)

type Interval struct {
	Left     int
	Right    int
	Priority int
}

type PriorityQueue []*Interval

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	interval := x.(*Interval)
	*pq = append(*pq, interval)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	interval := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return interval
}

func minGroups(intervals [][]int) int {
	left, right := &PriorityQueue{}, &PriorityQueue{}
	heap.Init(left)
	heap.Init(right)

	for _, interval := range intervals {
		obj := &Interval{
			Left:     interval[0],
			Right:    interval[1],
			Priority: interval[0],
		}
		heap.Push(left, obj)
	}

	maxIntersect := 0
	for left.Len() > 0 {
		intervalLeft := heap.Pop(left).(*Interval)
		intervalLeft.Priority = intervalLeft.Right
		heap.Push(right, intervalLeft)

		for right.Len() > 0 {
			intervalRight := heap.Pop(right).(*Interval)
			if intervalRight.Right >= intervalLeft.Left {
				heap.Push(right, intervalRight)
				break
			}
		}

		maxIntersect = max(maxIntersect, right.Len())
	}

	return maxIntersect
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
