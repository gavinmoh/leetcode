package main

import "container/heap"

type Element struct {
	Number int
	Index  int
}

type PriorityQueue []*Element

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Number == pq[j].Number {
		return pq[i].Index < pq[j].Index
	}
	return pq[i].Number < pq[j].Number
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	element := x.(*Element)
	*pq = append(*pq, element)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	element := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return element
}

func findScore(nums []int) int64 {
	marked := map[int]bool{}
	elements := &PriorityQueue{}
	heap.Init(elements)
	for index, num := range nums {
		heap.Push(elements, &Element{
			Number: num,
			Index:  index,
		})
	}

	score := int64(0)
	for elements.Len() > 0 {
		element := heap.Pop(elements).(*Element)
		// check if it's marked
		if _, ok := marked[element.Index]; ok {
			continue
		}
		marked[element.Index] = true
		marked[element.Index+1] = true
		marked[element.Index-1] = true

		score += int64(element.Number)
	}

	return score
}
