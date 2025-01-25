package main

import (
	"container/heap"
	"sort"
)

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

func lexicographicallySmallestArray(nums []int, limit int) []int {
	// create a new copy of nums
	// and sort it
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	// mapping of num into group
	numsGroup := map[int]int{}
	groupHeap := map[int]*IntHeap{}
	groupIdx := 0

	// initialize map
	numsGroup[sorted[0]] = groupIdx
	groupHeap[groupIdx] = &IntHeap{sorted[0]}
	heap.Init(groupHeap[groupIdx])

	for i := 1; i < len(sorted); i++ {
		if sorted[i]-sorted[i-1] > limit {
			groupIdx++
			groupHeap[groupIdx] = &IntHeap{}
			heap.Init(groupHeap[groupIdx])
		}
		numsGroup[sorted[i]] = groupIdx
		heap.Push(groupHeap[groupIdx], sorted[i])
	}

	for i := 0; i < len(nums); i++ {
		group := numsGroup[nums[i]]
		nums[i] = heap.Pop(groupHeap[group]).(int)
	}

	return nums
}
