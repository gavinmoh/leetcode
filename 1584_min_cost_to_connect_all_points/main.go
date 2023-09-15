package main

import "sort"

func minCostConnectPoints(points [][]int) int {
	// get number of points
	n := len(points)

	// calculate all edges between points and their distance
	edges := [][]int{}
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			x := points[i]
			y := points[j]
			distance := calculateDistanceBetween(x, y)
			edges = append(edges, []int{distance, i, j})
		}
	}

	// sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	// initialize variables for union-find algorithm
	parents := make([]int, n)
	for i := range parents {
		parents[i] = i
	}

	// union-find algorithm
	var unionFind func(x int) int
	unionFind = func(x int) int {
		if parents[x] != x {
			parents[x] = unionFind(parents[x])
		}
		return parents[x]
	}

	// find minimum spanning tree
	cost := 0
	for _, edge := range edges {
		weight := edge[0]
		x := edge[1]
		y := edge[2]
		if unionFind(x) == unionFind(y) {
			continue
		}
		parents[unionFind(x)] = unionFind(y)
		cost += weight
		n--
		if n == 1 {
			return cost
		}
	}
	return cost
}

// calculate manhattan distance between two points
func calculateDistanceBetween(x []int, y []int) int {
	xDistance := x[0] - y[0]
	yDistance := x[1] - y[1]

	if xDistance < 0 {
		xDistance = -xDistance
	}

	if yDistance < 0 {
		yDistance = -yDistance
	}

	return xDistance + yDistance
}
