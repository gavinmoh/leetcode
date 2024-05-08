package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findRelativeRanks(score []int) []string {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range score {
		heap.Push(h, -num)
	}

	scoreRank := make(map[int]int) // score -> rank
	for i := 1; h.Len() > 0; i++ {
		currentScore := -heap.Pop(h).(int)
		scoreRank[currentScore] = i
	}

	result := []string{}
	for _, currentScore := range score {
		var place string
		switch rank := scoreRank[currentScore]; rank {
		case 1:
			place = "Gold Medal"
		case 2:
			place = "Silver Medal"
		case 3:
			place = "Bronze Medal"
		default:
			place = fmt.Sprintf("%d", rank)
		}
		result = append(result, place)
	}

	return result
}
