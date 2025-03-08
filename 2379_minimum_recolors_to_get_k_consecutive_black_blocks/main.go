package main

func minimumRecolors(blocks string, k int) int {
	wCount := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			wCount++
		}
	}
	minCount := wCount

	if wCount == 0 {
		return 0
	}

	for i := k; i < len(blocks); i++ {
		if blocks[i-k] == 'W' {
			wCount--
		}

		if blocks[i] == 'W' {
			wCount++
		}

		if wCount < minCount {
			minCount = wCount
		}
	}

	return minCount
}
