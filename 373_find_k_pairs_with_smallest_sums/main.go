package main

import "container/heap"

type Pair struct {
	Num1 int
	Num2 int
	Sum  int
}

type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Sum < pq[j].Sum
}

func (pq *PriorityQueue) Push(x any) {
	pair := x.(*Pair)
	*pq = append(*pq, pair)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	pair := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return pair
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := &PriorityQueue{
		&Pair{Num1: 0, Num2: 0, Sum: nums1[0] + nums2[0]},
	}
	heap.Init(pq)

	m, n := len(nums1), len(nums2)
	result := [][]int{}

	visited := make(map[int]map[int]bool)
	visited[0] = map[int]bool{0: true}

	for k > 0 && pq.Len() > 0 {
		pair := heap.Pop(pq).(*Pair)
		i, j := pair.Num1, pair.Num2
		result = append(result, []int{nums1[i], nums2[j]})

		if i+1 < m {
			if _, ok := visited[i+1]; !ok {
				visited[i+1] = make(map[int]bool)
			}
			if _, ok := visited[i+1][j]; !ok {
				heap.Push(pq, &Pair{
					Num1: i + 1,
					Num2: j,
					Sum:  nums1[i+1] + nums2[j],
				})
				visited[i+1][j] = true
			}
		}

		if j+1 < n {
			if _, ok := visited[i]; !ok {
				visited[i] = make(map[int]bool)
			}
			if _, ok := visited[i][j+1]; !ok {
				heap.Push(pq, &Pair{
					Num1: i,
					Num2: j + 1,
					Sum:  nums1[i] + nums2[j+1],
				})
				visited[i][j+1] = true
			}
		}

		k--
	}

	return result
}
