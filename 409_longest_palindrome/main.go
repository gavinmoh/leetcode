package main

func longestPalindrome(s string) int {
	letters := make(map[rune]int)
	for _, letter := range s {
		letters[letter]++
	}
	length := 0
	hasOddFreq := false
	for _, freq := range letters {
		if freq%2 == 0 {
			length += freq
		} else {
			length += freq - 1
			hasOddFreq = true
		}
	}

	if hasOddFreq {
		return length + 1
	}

	return length
}
