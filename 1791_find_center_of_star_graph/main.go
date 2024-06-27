package main

func findCenter(edges [][]int) int {
	edge1, edge2 := edges[0], edges[1]

	if edge1[0] == edge2[0] || edge1[0] == edge2[1] {
		return edge1[0]
	}

	return edge1[1]
}
