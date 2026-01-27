package main

import "container/heap"

func minCost(n int, edges [][]int) int {
	adjacencyGraph := make([][][]int, n)
	for i := range adjacencyGraph {
		adjacencyGraph[i] = [][]int{}
	}

	for _, edge := range edges {
		from, to, cost := edge[0], edge[1], edge[2]
		adjacencyGraph[from] = append(adjacencyGraph[from], []int{to, cost})
		adjacencyGraph[to] = append(adjacencyGraph[to], []int{from, 2 * cost})
	}

	queue := &PriorityQueue{}
	heap.Init(queue)
	// start from node 0
	for _, neighbour := range adjacencyGraph[0] {
		node, cost := neighbour[0], neighbour[1]
		heap.Push(queue, &Path{
			Node: node,
			Cost: cost,
		})
	}

	visited := make([]bool, n)
	visited[0] = true

	for queue.Len() > 0 {
		path := heap.Pop(queue).(*Path)
		currentNode, currentCost := path.Node, path.Cost
		if currentNode == n-1 {
			return currentCost
		}

		if visited[currentNode] {
			continue
		}
		visited[currentNode] = true

		for _, neighbour := range adjacencyGraph[currentNode] {
			node, cost := neighbour[0], neighbour[1]
			heap.Push(queue, &Path{
				Node: node,
				Cost: currentCost + cost,
			})
		}
	}

	return -1
}

type Path struct {
	Node int
	Cost int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Path

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Path)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	path := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	path.index = -1 // for safety
	*pq = old[0 : n-1]
	return path
}
