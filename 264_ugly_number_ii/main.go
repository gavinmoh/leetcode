package main

import "container/heap"

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func nthUglyNumber(n int) int {
	uglyMap := map[int]bool{}
	uglyHeap := &IntHeap{}
	heap.Init(uglyHeap)
	uglyNumber := 1
	n--

	for n > 0 {
		if _, ok := uglyMap[uglyNumber*2]; !ok {
			uglyMap[uglyNumber*2] = true
			heap.Push(uglyHeap, uglyNumber*2)
		}
		if _, ok := uglyMap[uglyNumber*3]; !ok {
			uglyMap[uglyNumber*3] = true
			heap.Push(uglyHeap, uglyNumber*3)
		}
		if _, ok := uglyMap[uglyNumber*5]; !ok {
			uglyMap[uglyNumber*5] = true
			heap.Push(uglyHeap, uglyNumber*5)
		}

		uglyNumber = heap.Pop(uglyHeap).(int)
		n--
	}

	return uglyNumber
}
