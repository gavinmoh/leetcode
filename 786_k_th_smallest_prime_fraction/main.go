package main

import "container/heap"

type Fraction struct {
	Value       float64
	Numerator   int
	Denominator int
}

type MinHeap []*Fraction

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Value < h[j].Value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Fraction))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			numerator := arr[i]
			denominator := arr[j]
			value := float64(numerator) / float64(denominator)
			heap.Push(minHeap, &Fraction{
				Value:       value,
				Numerator:   numerator,
				Denominator: denominator,
			})
		}
	}

	var fraction *Fraction
	for ; k > 0; k-- {
		fraction = heap.Pop(minHeap).(*Fraction)
	}

	return []int{fraction.Numerator, fraction.Denominator}
}
