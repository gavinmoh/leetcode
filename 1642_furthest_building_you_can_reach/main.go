package main

import (
	"container/heap"
	"sort"
)

type MinHeap struct{ sort.IntSlice }

func (h *MinHeap) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }

func (h *MinHeap) Pop() any {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[:n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	maxHeap := &MinHeap{}

	for i := 0; i < n-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff <= 0 {
			continue
		}

		bricks -= diff
		heap.Push(maxHeap, -diff)

		if bricks < 0 {
			if ladders == 0 {
				return i
			}
			ladders -= 1
			bricks -= heap.Pop(maxHeap).(int)
		}
	}

	return n - 1
}
