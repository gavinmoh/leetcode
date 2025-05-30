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

func pickGifts(gifts []int, k int) int64 {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)

	for _, gift := range gifts {
		heap.Push(maxHeap, gift)
	}

	for i := 0; i < k; i++ {
		gift := heap.Pop(maxHeap).(int)
		sqrt := math.Sqrt(float64(gift))
		remain := int(sqrt)
		heap.Push(maxHeap, remain)
	}

	total := int64(0)
	for maxHeap.Len() > 0 {
		gift := heap.Pop(maxHeap).(int)
		total += int64(gift)
	}

	return total
}
