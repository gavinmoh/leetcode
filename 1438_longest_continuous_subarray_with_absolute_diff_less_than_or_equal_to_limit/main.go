package main

import (
	"container/heap"
)

type Item struct {
	value int
	index int
}

type PriorityQueue []Item

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].value < h[j].value }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func longestSubarray(nums []int, limit int) int {
	maxHeap := &PriorityQueue{}
	minHeap := &PriorityQueue{}
	heap.Init(maxHeap)
	heap.Init(minHeap)

	left := 0
	maxLength := 0

	for right, num := range nums {
		heap.Push(maxHeap, Item{value: -num, index: right})
		heap.Push(minHeap, Item{value: num, index: right})

		for -(*maxHeap)[0].value-(*minHeap)[0].value > limit {
			left = min((*maxHeap)[0].index, (*minHeap)[0].index) + 1

			for (*maxHeap)[0].index < left {
				heap.Pop(maxHeap)
			}
			for (*minHeap)[0].index < left {
				heap.Pop(minHeap)
			}
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
