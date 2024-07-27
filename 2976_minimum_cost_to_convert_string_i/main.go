package main

import "container/heap"

type Edge struct {
	Node     byte
	Distance int
}

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

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	adj := make(map[byte][]*Edge)
	for i := 0; i < len(original); i++ {
		edge := &Edge{
			Node:     changed[i],
			Distance: cost[i],
		}
		if _, ok := adj[original[i]]; !ok {
			adj[original[i]] = []*Edge{}
		}
		adj[original[i]] = append(adj[original[i]], edge)
	}

	var dijkstra func(src, dest byte) int
	dijkstra = func(src, dest byte) int {
		minHeap := MinHeap{}
		heap.Init(&minHeap)
		heap.Push(&minHeap, &Edge{
			Node:     src,
			Distance: 0,
		})
		visited := make(map[byte]bool)

		for minHeap.Len() > 0 {
			edge := heap.Pop(&minHeap).(*Edge)

			if edge.Node == dest {
				return edge.Distance
			}

			if _, ok := visited[edge.Node]; ok {
				continue
			}

			visited[edge.Node] = true
			for _, neighbour := range adj[edge.Node] {
				neighbourDistance := neighbour.Distance + edge.Distance
				heap.Push(&minHeap, &Edge{
					Node:     neighbour.Node,
					Distance: neighbourDistance,
				})
			}
		}

		return -1
	}

	totalCost := int64(0)
	cache := make(map[byte]map[byte]int64)
	for i := 0; i < len(source); i++ {
		src, dest := source[i], target[i]
		if _, ok := cache[src]; !ok {
			cache[src] = make(map[byte]int64)
		}
		if cachedResult, ok := cache[src][dest]; ok {
			totalCost += cachedResult
			continue
		}

		cost := int64(dijkstra(src, dest))
		if cost == -1 {
			return -1
		}

		cache[src][dest] = cost
		totalCost += cost
	}

	return totalCost
}
