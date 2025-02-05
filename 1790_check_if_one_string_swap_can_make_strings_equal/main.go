package main

func areAlmostEqual(s1 string, s2 string) bool {
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)

	posDiff := 0
	for i := 0; i < len(s1); i++ {
		c1, c2 := s1[i], s2[i]
		if c1 != c2 {
			posDiff++

			if posDiff > 2 {
				return false
			}
		}

		freq1[c1-'a']++
		freq2[c2-'a']++
	}

	freqDiff := 0
	for i := 0; i < 26; i++ {
		c1Count, c2Count := freq1[i], freq2[i]
		freqDiff += abs(c1Count - c2Count)

		if freqDiff > 1 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
