package main

import "sort"

func minimumEffortPath(heights [][]int) int {
	rowCount := len(heights)
	colCount := len(heights[0])

	// compute the edges
	edges := make([][]int, 0)
	for i, row := range heights {
		for j, cell := range row {
			if i < rowCount-1 { // 2 row before hitting bottom
				cellBelow := heights[i+1][j]
				cost := absolute(cell - cellBelow) // cost of going down
				edge := []int{cost, i*colCount + j, (i+1)*colCount + j}
				edges = append(edges, edge)
			}
			if j < colCount-1 { // 2 col before hitting right
				cellRight := heights[i][j+1]
				cost := absolute(cell - cellRight) // cost of going right
				edge := []int{cost, i*colCount + j, i*colCount + j + 1}
				edges = append(edges, edge)
			}
		}
	}

	// sort the edges by cost
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	// initialize variables for union-find algorithm
	parents := make([]int, rowCount*colCount)
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

	for _, edge := range edges {
		cost, i, j := edge[0], edge[1], edge[2]
		parents[unionFind(i)] = unionFind(j)
		if unionFind(0) == unionFind(rowCount*colCount-1) {
			return cost
		}
	}
	return 0
}

func absolute(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
