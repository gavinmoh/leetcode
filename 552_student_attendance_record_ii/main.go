package main

func checkRecord(n int) int {
	modulo := 1_000_000_007
	cache := make([][][]int, n+1)
	for i := range cache {
		cache[i] = make([][]int, 2)
		for j := range cache[i] {
			cache[i][j] = make([]int, 3)
			for k := range cache[i][j] {
				cache[i][j][k] = -1
			}
		}
	}

	var dfs func(m int, absentCount int, consecutiveLateCount int) int
	dfs = func(m int, absentCount int, consecutiveLateCount int) int {
		if absentCount >= 2 || consecutiveLateCount >= 3 {
			return 0
		}

		if m == 0 {
			return 1
		}

		if cache[m][absentCount][consecutiveLateCount] != -1 {
			return cache[m][absentCount][consecutiveLateCount]
		}

		absent := dfs(m-1, absentCount+1, 0)
		late := dfs(m-1, absentCount, consecutiveLateCount+1)
		present := dfs(m-1, absentCount, 0)
		result := (absent + late + present) % modulo
		cache[m][absentCount][consecutiveLateCount] = result

		return result
	}

	return dfs(n, 0, 0)
}
