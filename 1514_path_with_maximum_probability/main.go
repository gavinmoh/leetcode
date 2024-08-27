package main

import (
	"container/heap"
	"math"
)

type Edge struct {
	Node        int
	Probability float64
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Probability > pq[j].Probability }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	// initialize adjacency list
	adj := make(map[int][]*Edge)
	for i := 0; i < n; i++ {
		adj[i] = []*Edge{}
	}

	for i, edge := range edges {
		probability := math.Log10(succProb[i])
		adj[edge[0]] = append(adj[edge[0]], &Edge{
			Node:        edge[1],
			Probability: probability,
		})
		adj[edge[1]] = append(adj[edge[1]], &Edge{
			Node:        edge[0],
			Probability: probability,
		})
	}

	queue := &PriorityQueue{}
	heap.Init(queue)
	for _, edge := range adj[start_node] {
		heap.Push(queue, edge)
	}

	visited := map[int]bool{start_node: true}
	for queue.Len() > 0 {
		edge := heap.Pop(queue).(*Edge)
		if _, ok := visited[edge.Node]; ok {
			continue
		}

		visited[edge.Node] = true
		if edge.Node == end_node {
			return math.Pow(10, edge.Probability)
		}

		for _, neighbour := range adj[edge.Node] {
			heap.Push(queue, &Edge{
				Node:        neighbour.Node,
				Probability: edge.Probability + neighbour.Probability,
			})
		}
	}

	return float64(0)
}
