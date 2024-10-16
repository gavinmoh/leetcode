package main

import "container/heap"

type Letter struct {
	Char       byte
	Occurrence int
}

type PriorityQueue []*Letter

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Occurrence > pq[j].Occurrence
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	letter := x.(*Letter)
	*pq = append(*pq, letter)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	letter := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return letter
}

func longestDiverseString(a int, b int, c int) string {
	letters := &PriorityQueue{}
	heap.Init(letters)
	if a > 0 {
		heap.Push(letters, &Letter{Char: 'a', Occurrence: a})
	}
	if b > 0 {
		heap.Push(letters, &Letter{Char: 'b', Occurrence: b})
	}
	if c > 0 {
		heap.Push(letters, &Letter{Char: 'c', Occurrence: c})
	}

	happy := ""
	for letters.Len() > 0 {
		first := heap.Pop(letters).(*Letter)
		firstAdded := false
		for i := 0; i < 2 && first.Occurrence > 0; i++ {
			if len(happy) > 1 && happy[len(happy)-2] == first.Char && happy[len(happy)-1] == first.Char {
				break
			}
			happy += string(first.Char)
			firstAdded = true
			first.Occurrence--
		}

		if letters.Len() > 0 {
			second := heap.Pop(letters).(*Letter)
			happy += string(second.Char)
			second.Occurrence--

			if second.Occurrence > 0 {
				heap.Push(letters, second)
			}
		}

		if !firstAdded {
			return happy
		}

		if first.Occurrence > 0 {
			heap.Push(letters, first)
		}
	}

	return happy
}
