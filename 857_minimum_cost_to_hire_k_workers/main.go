package main

import (
	"container/heap"
	"math"
)

type Worker struct {
	Quality float64
	Wage    float64
	Ratio   float64
}

type MinHeap []*Worker

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Ratio < h[j].Ratio }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Worker))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []*Worker

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Quality > h[j].Quality }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Worker))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	totalCost := math.MaxFloat64
	currentTotalQuality := float64(0)
	ratioHeap, qualityHeap := &MinHeap{}, &MaxHeap{}
	heap.Init(ratioHeap)
	heap.Init(qualityHeap)

	for i := 0; i < n; i++ {
		workerQuality := float64(quality[i])
		workerWage := float64(wage[i])
		worker := &Worker{
			Quality: workerQuality,
			Wage:    workerWage,
			Ratio:   workerWage / workerQuality,
		}
		heap.Push(ratioHeap, worker)
	}

	for ratioHeap.Len() > 0 {
		currentWorker := heap.Pop(ratioHeap).(*Worker)
		heap.Push(qualityHeap, currentWorker)
		currentTotalQuality += currentWorker.Quality
		if qualityHeap.Len() > k {
			poppedWorker := heap.Pop(qualityHeap).(*Worker)
			currentTotalQuality -= poppedWorker.Quality
		}
		if qualityHeap.Len() == k {
			currentCost := currentTotalQuality * currentWorker.Ratio
			if currentCost < totalCost {
				totalCost = currentCost
			}
		}
	}
	return totalCost
}
