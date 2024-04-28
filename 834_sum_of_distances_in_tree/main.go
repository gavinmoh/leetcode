package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	nodes := make([][]int, n)
	for _, edge := range edges {
		left, right := edge[0], edge[1]
		nodes[left] = append(nodes[left], right)
		nodes[right] = append(nodes[right], left)
	}
	subtrees := make([]int, n)
	visited1 := make([]bool, n)
	var dfs1 func(i int) int
	dfs1 = func(i int) int {
		sum := 1
		for _, node := range nodes[i] {
			if visited1[node] {
				continue
			}
			visited1[node] = true
			sum += dfs1(node)
		}
		subtrees[i] = sum
		return sum
	}

	visited2 := make([]bool, n)
	var dfs2 func(i int) int
	dfs2 = func(i int) int {
		sum := 0
		for _, node := range nodes[i] {
			if visited2[node] {
				continue
			}
			visited2[node] = true
			sum += dfs2(node)
		}
		sum += subtrees[i] - 1
		return sum
	}

	result := make([]int, n)
	visited3 := make([]bool, n)
	var dfs3 func(i int)
	dfs3 = func(i int) {
		for _, node := range nodes[i] {
			if visited3[node] {
				continue
			}
			visited3[node] = true
			b := subtrees[node]
			a := n - b
			result[node] = result[i] + a - b
			dfs3(node)
		}
	}

	visited1[0] = true
	dfs1(0)
	visited2[0] = true
	result[0] = dfs2(0)
	visited3[0] = true
	dfs3(0)

	return result
}
