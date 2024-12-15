package main

import "container/heap"

type Class struct {
	Students int
	Passes   int
}

func (c *Class) Gain() float64 {
	return float64(c.Passes+1)/float64(c.Students+1) - c.Ratio()
}

func (c *Class) Ratio() float64 {
	return float64(c.Passes) / float64(c.Students)
}

type PriorityQueue []*Class

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Gain() > pq[j].Gain() }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	class := x.(*Class)
	*pq = append(*pq, class)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	class := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually
	*pq = old[0 : n-1]
	return class
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for _, class := range classes {
		passes, students := class[0], class[1]
		heap.Push(pq, &Class{
			Students: students,
			Passes:   passes,
		})
	}

	for i := 0; i < extraStudents; i++ {
		class := heap.Pop(pq).(*Class)
		class.Students++
		class.Passes++
		heap.Push(pq, class)
	}

	totalRatio := float64(0)
	for pq.Len() > 0 {
		class := heap.Pop(pq).(*Class)
		totalRatio += class.Ratio()
	}

	return totalRatio / float64(len(classes))
}
