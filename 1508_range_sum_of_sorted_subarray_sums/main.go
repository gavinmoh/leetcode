package main

import "container/heap"

const MOD = int(1e9) + 7

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

func rangeSum(nums []int, n int, left int, right int) int {
	sums := &IntHeap{}
	heap.Init(sums)
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			sum := 0
			for _, num := range nums[i:j] {
				sum += num
			}
			heap.Push(sums, sum)
		}
	}

	result := 0
	for i := 1; i <= right; i++ {
		num := heap.Pop(sums).(int)
		if i >= left {
			result += num
			result %= MOD
		}
	}

	return result
}
