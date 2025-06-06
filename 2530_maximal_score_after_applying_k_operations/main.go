package main

import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxKelements(nums []int, k int) int64 {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)
	for _, num := range nums {
		heap.Push(maxHeap, num)
	}

	result := int64(0)
	for i := 0; i < k; i++ {
		num := heap.Pop(maxHeap).(int)
		result += int64(num)
		newNum := math.Ceil(float64(num) / float64(3))
		heap.Push(maxHeap, int(newNum))
	}

	return result
}
