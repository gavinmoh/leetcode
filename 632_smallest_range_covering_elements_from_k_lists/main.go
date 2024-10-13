package main

import "container/heap"

type Item struct {
	List  int
	Num   int
	Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Num < pq[j].Num
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

func smallestRange(nums [][]int) []int {
	left, right := nums[0][0], nums[0][0]
	pq := &PriorityQueue{}
	heap.Init(pq)
	for i, list := range nums {
		left = min(left, list[0])
		right = max(right, list[0])
		heap.Push(pq, &Item{
			List:  i, // which list
			Num:   list[0],
			Index: 0, // index of number in the list
		})
	}

	result := []int{left, right}
	for {
		item := heap.Pop(pq).(*Item)
		index := item.Index + 1
		if index == len(nums[item.List]) {
			return result
		}

		num := nums[item.List][index]
		heap.Push(pq, &Item{
			List:  item.List,
			Num:   num,
			Index: index,
		})
		right = max(right, num)
		left = (*pq)[0].Num
		if (right - left) < (result[1] - result[0]) {
			result = []int{left, right}
		}
	}
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
