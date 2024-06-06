package main

import "container/heap"

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

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	cards := &IntHeap{}
	heap.Init(cards)
	for _, card := range hand {
		heap.Push(cards, card)
	}

	for i := 0; i < len(hand)/groupSize; i++ {
		temp := []int{}
		currentCard := heap.Pop(cards).(int)
		for j := 1; j < groupSize; j++ {
			if cards.Len() == 0 {
				return false
			}
			nextCard := heap.Pop(cards).(int)
			for nextCard == currentCard && cards.Len() > 0 {
				temp = append(temp, nextCard)
				nextCard = heap.Pop(cards).(int)
			}
			if nextCard != currentCard+1 {
				return false
			}
			currentCard = nextCard
		}
		for _, card := range temp {
			heap.Push(cards, card)
		}
	}

	return true
}
