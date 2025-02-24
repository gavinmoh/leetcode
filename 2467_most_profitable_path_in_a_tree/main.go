package main

import (
	"math"
)

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(edges)
	adjList := map[int][]int{}
	for i := 0; i < n+1; i++ {
		adjList[i] = []int{}
	}

	for _, edge := range edges {
		left, right := edge[0], edge[1]
		adjList[left] = append(adjList[left], right)
		adjList[right] = append(adjList[right], left)
	}

	// find out the path bob taken
	bobVisited := map[int]bool{}
	var dfs func(node int, path []int) ([]int, bool)
	dfs = func(node int, path []int) ([]int, bool) {
		if _, ok := bobVisited[node]; ok {
			return []int{}, false
		}
		bobVisited[node] = true
		path = append(path, node)

		if node == 0 {
			return path, true
		}

		for _, neighbour := range adjList[node] {
			if p, found := dfs(neighbour, path); found {
				return p, found
			}
		}

		bobVisited[node] = false

		return []int{}, false
	}

	bobPath, _ := dfs(bob, []int{})
	bobTime := map[int]int{}
	for time, node := range bobPath {
		bobTime[node] = time
	}

	maxIncome := math.MinInt
	aliceVisited := map[int]bool{}
	var aliceDFS func(node, time, income int)
	aliceDFS = func(node, time, income int) {
		aliceVisited[node] = true

		reward := amount[node]
		if bTime, ok := bobTime[node]; ok {
			if bTime == time {
				reward /= 2
			} else if bobTime[node] < time {
				reward = 0
			}
		}

		income += reward
		isLeafNode := true

		for _, neighbour := range adjList[node] {
			if _, ok := aliceVisited[neighbour]; !ok {
				aliceDFS(neighbour, time+1, income)
				isLeafNode = false
			}
		}

		if isLeafNode && income > maxIncome {
			maxIncome = income
		}

		aliceVisited[node] = false
	}

	aliceDFS(0, 0, 0)

	return maxIncome
}
