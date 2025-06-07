package main

import (
	"container/heap"
	"sort"
	"strings"
)

type PriorityQueue [][]int // [][]int{{char, index}}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i][0] == pq[j][0] {
		return pq[i][1] > pq[j][1]
	}
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.([]int)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func clearStars(s string) string {
	queue := &PriorityQueue{}
	heap.Init(queue)

	for i, c := range s {
		if c == '*' && queue.Len() > 0 {
			heap.Pop(queue)
		} else {
			heap.Push(queue, []int{int(c - 'a'), i})
		}
	}

	chars := [][]int{}
	for queue.Len() > 0 {
		item := heap.Pop(queue).([]int)
		chars = append(chars, item)
	}

	sort.SliceStable(chars, func(i, j int) bool {
		return chars[i][1] < chars[j][1]
	})

	var result strings.Builder
	for _, char := range chars {
		result.WriteRune('a' + rune(char[0]))
	}

	return result.String()
}
