package main

import "math"

func minimumCost(n int, edges [][]int, query [][]int) []int {
	// Create the adjacency list of the graph
	adjList := make([][]struct{ node, weight int }, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], struct{ node, weight int }{edge[1], edge[2]})
		adjList[edge[1]] = append(adjList[edge[1]], struct{ node, weight int }{edge[0], edge[2]})
	}

	visited := make([]bool, n)
	components := make([]int, n)
	var componentCost []int
	componentId := 0

	// Perform DFS for each unvisited node to identify components and calculate their costs
	for node := 0; node < n; node++ {
		if !visited[node] {
			// Get the component cost and mark all nodes in the component
			cost := getComponentCost(node, adjList, visited, components, componentId)
			componentCost = append(componentCost, cost)
			componentId++
		}
	}

	var answer []int
	for _, q := range query {
		start, end := q[0], q[1]
		if components[start] == components[end] {
			answer = append(answer, componentCost[components[start]])
		} else {
			answer = append(answer, -1)
		}
	}
	return answer
}

// Helper function to calculate the cost of a component using DFS
func getComponentCost(node int, adjList [][]struct{ node, weight int }, visited []bool, components []int, componentId int) int {
	currentCost := math.MaxInt32 // Equivalent to INT_MAX
	components[node] = componentId
	visited[node] = true

	for _, neighbor := range adjList[node] {
		currentCost &= neighbor.weight
		if !visited[neighbor.node] {
			currentCost &= getComponentCost(neighbor.node, adjList, visited, components, componentId)
		}
	}
	return currentCost
}
