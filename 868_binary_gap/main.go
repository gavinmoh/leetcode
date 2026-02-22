package main

func binaryGap(n int) int {
	longest := 0
	prevOne := -1

	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if prevOne != -1 {
				distance := i - prevOne
				if distance > longest {
					longest = distance
				}
			}
			prevOne = i
		}
		n >>= 1
	}

	return longest
}
