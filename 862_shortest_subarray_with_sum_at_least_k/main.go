package main

import (
	"container/heap"
	"math"
)

type Item struct {
	Index int
	Sum   int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Sum < pq[j].Sum
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

func shortestSubarray(nums []int, k int) int {
	minLength := math.MaxInt
	sum := 0
	pq := &PriorityQueue{}
	heap.Init(pq)

	for i, num := range nums {
		sum += num
		if sum >= k {
			minLength = min(minLength, i+1)
		}
		for pq.Len() > 0 {
			item := heap.Pop(pq).(*Item)
			if sum-item.Sum >= k {
				minLength = min(minLength, i-item.Index)
			} else {
				heap.Push(pq, item)
				break
			}
		}
		heap.Push(pq, &Item{
			Index: i,
			Sum:   sum,
		})
	}

	if minLength == math.MaxInt {
		return -1
	}

	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
