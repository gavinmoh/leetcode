package main

import "slices"

func validArrangement(pairs [][]int) [][]int {
	adjacencyMatrix := map[int][]int{}
	inDegree, outDegree := map[int]int{}, map[int]int{}

	for _, pair := range pairs {
		start, end := pair[0], pair[1]
		if _, ok := adjacencyMatrix[start]; !ok {
			adjacencyMatrix[start] = []int{}
		}
		adjacencyMatrix[start] = append(adjacencyMatrix[start], end)
		outDegree[start]++
		inDegree[end]++
	}

	result := []int{}

	// find start node
	startNode := -1
	for node, count := range outDegree {
		if count == inDegree[node]+1 {
			startNode = node
			break
		}
	}

	if startNode == -1 {
		startNode = pairs[0][0]
	}

	nodeStack := []int{startNode}

	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		if len(adjacencyMatrix[node]) > 0 {
			// visit next node
			nextNode := adjacencyMatrix[node][0]
			adjacencyMatrix[node] = adjacencyMatrix[node][1:]
			nodeStack = append(nodeStack, nextNode)
		} else {
			// add node to result when no more neighbours to visit
			result = append(result, node)
			nodeStack = nodeStack[0 : len(nodeStack)-1]
		}
	}

	slices.Reverse(result)

	// reconstructing the pairs
	pairedResult := [][]int{}
	for i := 1; i < len(result); i++ {
		pairedResult = append(pairedResult, []int{result[i-1], result[i]})
	}

	return pairedResult
}
