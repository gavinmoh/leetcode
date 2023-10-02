package main

func winnerOfGame(colors string) bool {
	aCount := 0
	bCount := 0
	for i := 0; i < len(colors)-2; i++ {
		if colors[i] == 'A' && colors[i+1] == 'A' && colors[i+2] == 'A' {
			aCount++
		}
		if colors[i] == 'B' && colors[i+1] == 'B' && colors[i+2] == 'B' {
			bCount++
		}
	}

	return aCount > bCount
}
