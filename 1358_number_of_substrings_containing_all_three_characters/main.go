package main

func numberOfSubstrings(s string) int {
	freq := map[byte]int{'a': 0, 'b': 0, 'c': 0}
	left, right := 0, 0
	result := 0

	for right < len(s) {
		c := s[right]
		freq[c]++

		for freq['a'] > 0 && freq['b'] > 0 && freq['c'] > 0 {
			result += len(s) - right

			leftChar := s[left]
			freq[leftChar]--
			left++
		}

		right++
	}

	return result
}
