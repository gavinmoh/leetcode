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

type NumberContainers struct {
	Indices map[int]*IntHeap
	Numbers map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		Indices: map[int]*IntHeap{},
		Numbers: map[int]int{},
	}
}

func (this *NumberContainers) Change(index int, number int) {
	this.Numbers[index] = number
	if _, ok := this.Indices[number]; !ok {
		minHeap := &IntHeap{}
		heap.Init(minHeap)
		this.Indices[number] = minHeap
	}
	heap.Push(this.Indices[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if minHeap, ok := this.Indices[number]; ok && minHeap.Len() > 0 {
		for minHeap.Len() > 0 {
			idx := heap.Pop(minHeap).(int)
			if this.Numbers[idx] == number {
				heap.Push(minHeap, idx) // push back
				return idx
			}
		}
	}

	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
