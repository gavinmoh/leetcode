package main

func countPalindromicSubsequence(s string) int {
	n := len(s)
	answer := 0
	// create a unique set of letters
	letters := make(map[rune]bool, n)
	for _, letter := range s {
		if !letters[letter] {
			letters[letter] = true
		}
	}

	for letter := range letters {
		var first, last int
		// find the first occurrence of letter in s
		for i := 0; i < n; i++ {
			if rune(s[i]) == letter {
				first = i
				break
			}
		}
		// find the last occurrence
		for i := n - 1; i >= 0; i-- {
			if rune(s[i]) == letter {
				last = i
				break
			}
		}

		// no palindrome can be formed
		if last-first < 2 {
			continue
		}

		substr := s[first+1 : last]
		between := make(map[rune]bool, last-first)
		for _, char := range substr {
			if !between[char] {
				answer++
				between[char] = true
			}
		}
	}

	return answer

}
