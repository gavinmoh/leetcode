package main

func minBitFlips(start int, goal int) int {
	diff := start ^ goal
	count := 0
	for diff > 0 {
		if (diff & 1) == 1 {
			count++
		}
		diff = diff >> 1
	}

	return count
}
