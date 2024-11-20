package main

func findChampion(n int, edges [][]int) int {
	nodes := map[int]bool{}
	for _, edge := range edges {
		nodes[edge[1]] = true
	}

	champion := -1
	for i := 0; i < n; i++ {
		if _, ok := nodes[i]; !ok {
			if champion != -1 {
				return -1
			}
			champion = i
		}
	}

	return champion
}
