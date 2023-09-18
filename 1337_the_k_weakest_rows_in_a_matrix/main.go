package main

func kWeakestRows(mat [][]int, k int) []int {
	totalSoldiers := make(map[int]int)
	totalMap := make(map[int][]int)

	for i, row := range mat {
		for _, spot := range row {
			if spot == 1 {
				totalSoldiers[i]++
			}
		}
		totalMap[totalSoldiers[i]] = append(totalMap[totalSoldiers[i]], i)
	}

	weakest := make([]int, 0)
	for i := 0; i <= len(mat[0]); i++ {
		if len(totalMap[i]) > 0 {
			for _, v := range totalMap[i] {
				if len(weakest) < k {
					weakest = append(weakest, v)
				}
			}
		}
	}

	return weakest

}
