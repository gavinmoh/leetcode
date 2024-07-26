package main

import "container/heap"

type Edge struct {
	Node     int
	Distance int
}

// An IntHeap is a min-heap of ints.
type MinHeap []*Edge

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Edge))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	adj := make(map[int][]*Edge)
	for _, edge := range edges {
		v1, v2, distance := edge[0], edge[1], edge[2]
		edge1 := &Edge{
			Node:     v2,
			Distance: distance,
		}
		edge2 := &Edge{
			Node:     v1,
			Distance: distance,
		}
		if _, ok := adj[v1]; !ok {
			adj[v1] = []*Edge{}
		}
		adj[v1] = append(adj[v1], edge1)
		if _, ok := adj[v2]; !ok {
			adj[v2] = []*Edge{}
		}
		adj[v2] = append(adj[v2], edge2)
	}

	var dijkstra func(source int) int
	dijkstra = func(source int) int {
		minHeap := MinHeap{}
		heap.Init(&minHeap)
		heap.Push(&minHeap, &Edge{
			Node:     source,
			Distance: 0,
		})
		visited := make(map[int]bool)

		for minHeap.Len() > 0 {
			edge := heap.Pop(&minHeap).(*Edge)

			if _, ok := visited[edge.Node]; ok {
				continue
			}

			visited[edge.Node] = true
			for _, neighbour := range adj[edge.Node] {
				neighbourDistance := neighbour.Distance + edge.Distance
				if neighbourDistance <= distanceThreshold {
					heap.Push(&minHeap, &Edge{
						Node:     neighbour.Node,
						Distance: neighbourDistance,
					})
				}
			}
		}

		return len(visited) - 1
	}

	result, minCount := -1, n
	for i := 0; i < n; i++ {
		count := dijkstra(i)
		if count <= minCount {
			result, minCount = i, count
		}
	}

	return result
}
