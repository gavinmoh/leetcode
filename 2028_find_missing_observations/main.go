package main

func missingRolls(rolls []int, mean int, n int) []int {
	mSum := 0
	for _, roll := range rolls {
		mSum += roll
	}
	totalSum := (n + len(rolls)) * mean
	nSum := totalSum - mSum

	if nSum > n*6 || nSum < n {
		return []int{}
	}

	nMean := nSum / n
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = nMean
		nSum -= nMean
	}

	for i := 0; nSum > 0; i++ {
		result[i]++
		nSum--
	}

	return result
}
