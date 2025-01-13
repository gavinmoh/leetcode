package main

func minimumLength(s string) int {
	freq := map[rune]int{}
	for _, char := range s {
		freq[char]++
	}

	operation := 0
	for _, count := range freq {
		if count < 3 {
			continue
		}

		if count%2 == 1 {
			operation += count - 1 // left one
		} else {
			operation += count - 2 // left two
		}
	}

	return len(s) - operation
}
