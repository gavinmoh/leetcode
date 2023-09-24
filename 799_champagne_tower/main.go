package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	previousRow := make([]float64, 1)
	previousRow[0] = float64(poured)
	for i := 0; i < query_row; i++ {
		currentRow := make([]float64, i+2)
		for j := 0; j < len(previousRow); j++ {
			if previousRow[j] > 1 {
				currentRow[j] += (previousRow[j] - 1) / 2
				currentRow[j+1] += (previousRow[j] - 1) / 2
			}
		}
		previousRow = currentRow
	}

	if previousRow[query_glass] > 1 {
		return 1
	}

	return previousRow[query_glass]

}
