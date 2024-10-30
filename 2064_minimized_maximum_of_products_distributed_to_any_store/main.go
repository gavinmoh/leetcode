package main

import (
	"container/heap"
	"math"
)

type Product struct {
	Quantity   float64
	StoreCount float64
	Ratio      float64
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Product

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Ratio > pq[j].Ratio
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	product := x.(*Product)
	*pq = append(*pq, product)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	product := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return product
}

func minimizedMaximum(n int, quantities []int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, quantity := range quantities {
		ratio := float64(quantity) / 1.0
		heap.Push(pq, &Product{
			Quantity:   float64(quantity),
			StoreCount: 1.0,
			Ratio:      ratio,
		})
		n--
	}

	for n > 0 {
		product := heap.Pop(pq).(*Product)
		product.StoreCount++
		product.Ratio = product.Quantity / product.StoreCount
		heap.Push(pq, product)
		n--
	}

	product := heap.Pop(pq).(*Product)

	return int(math.Ceil(product.Ratio))
}
