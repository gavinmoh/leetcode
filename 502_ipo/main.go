package main

import (
	"container/heap"
	"sort"
)

type Project struct {
	Profit  int
	Capital int
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := []Project{}
	for i := 0; i < len(profits); i++ {
		project := Project{Profit: profits[i], Capital: capital[i]}
		projects = append(projects, project)
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].Capital < projects[j].Capital
	})

	maxHeap := IntHeap{}
	heap.Init(&maxHeap)

	for i := 0; k > 0; {
		for i < len(profits) && projects[i].Capital <= w {
			heap.Push(&maxHeap, projects[i].Profit)
			i++
		}

		if maxHeap.Len() == 0 {
			break
		}

		w += heap.Pop(&maxHeap).(int)
		k--
	}

	return w
}
