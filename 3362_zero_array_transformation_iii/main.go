package main

import (
	"container/heap"
	"sort"
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

func (h *IntHeap) Top() any {
	arr := *h
	return arr[0]
}

func maxRemoval(nums []int, queries [][]int) int {
	sort.SliceStable(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	intHeap := &IntHeap{}
	heap.Init(intHeap)
	diffArray := make([]int, len(nums)+1)
	count := 0

	for i, j := 0, 0; i < len(nums); i++ {
		count += diffArray[i]
		for j < len(queries) && queries[j][0] == i {
			heap.Push(intHeap, queries[j][1])
			j++
		}

		for count < nums[i] && intHeap.Len() > 0 && intHeap.Top().(int) >= i {
			count += 1
			diffArray[heap.Pop(intHeap).(int)+1] -= 1
		}

		if count < nums[i] {
			return -1
		}
	}

	return intHeap.Len()
}
