package main

import (
	"container/heap"
	"strings"
)

type Item struct {
	Char  rune
	Count int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Char > pq[j].Char
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

func repeatLimitedString(s string, repeatLimit int) string {
	charMap := map[rune]int{}
	for _, char := range s {
		charMap[char]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	for char, count := range charMap {
		heap.Push(pq, &Item{
			Char:  char,
			Count: count,
		})
	}

	var result strings.Builder
	var prev *Item
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		limit := min(repeatLimit, item.Count)
		if prev != nil && prev.Count > 0 {
			heap.Push(pq, prev)
			if prev.Char > item.Char {
				limit = 1
			}
		}

		for i := limit; i > 0; i-- {
			result.WriteRune(item.Char)
			item.Count--
		}
		prev = item
	}

	return result.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
