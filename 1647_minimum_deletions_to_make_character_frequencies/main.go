package main

func MinDeletions(s string) int {
	// count the frequency of each characters
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	// find the characters that have the same frequency
	freqCount := make(map[int]int) // key: frequency, value: count
	for _, freq := range charCount {
		freqCount[freq]++
	}

	deletions := 0
	sorted := false
	for !sorted {
		sorted = true
		for freq, count := range freqCount {
			for count > 1 {
				sorted = false
				deletions++       // delete one character
				count--           // decrease the count of the character
				freqCount[freq]-- // decrease the count of the frequency
				if freq > 1 {     // we do not care if the frequency is 0
					freqCount[freq-1]++ // increase the count of the frequency - 1
				}
			}
		}
	}

	return deletions
}
