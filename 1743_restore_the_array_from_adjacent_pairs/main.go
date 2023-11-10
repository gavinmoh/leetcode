package main

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)

	if n == 1 {
		return adjacentPairs[0]
	}

	frequencies := make(map[int]int, n+1)
	pairs := make(map[int][][]int, n)
	for _, pair := range adjacentPairs {
		left, right := pair[0], pair[1]
		pairs[left] = append(pairs[left], pair)
		pairs[right] = append(pairs[right], pair)

		frequencies[left] += 1
		frequencies[right] += 1
	}

	var first int
	for i, freq := range frequencies {
		if freq == 1 {
			first = i
			break
		}
	}

	visited := make(map[int][][]int, n+1)
	var dfs func([]int, int) []int
	dfs = func(reconstructed []int, val int) []int {
		reconstructed = append(reconstructed, val)
		if len(reconstructed) == n+1 {
			return reconstructed
		}
		var next int

		vis := visited[val]
		for _, pair := range pairs[val] {
			if !contains(vis, pair) {
				left, right := pair[0], pair[1]
				if left == val {
					next = right
				} else {
					next = left
				}
				visited[left] = append(visited[left], pair)
				visited[right] = append(visited[right], pair)
				break
			}
		}
		return dfs(reconstructed, next)
	}

	return dfs([]int{}, first)
}

func contains(arrays [][]int, array []int) bool {
	for _, arr := range arrays {
		if arr[0] == array[0] && arr[1] == array[1] {
			return true
		}
	}
	return false
}
