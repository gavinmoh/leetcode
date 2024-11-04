package main

func largestCombination(candidates []int) int {
	bitsCount := make([]int, 24)

	for _, num := range candidates {
		for position := 0; num > 0; position++ {
			if (num & 1) == 1 {
				bitsCount[position]++
			}
			num >>= 1
		}
	}
	maxCount := 0
	for _, count := range bitsCount {
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
