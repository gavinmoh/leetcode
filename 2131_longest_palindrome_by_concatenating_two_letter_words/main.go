package main

func longestPalindrome(words []string) int {
	freq := map[string]int{}
	for _, word := range words {
		freq[word]++
	}

	palindrome := 0
	center := 0
	for word, count := range freq {
		if isPalindrome(word) {
			// if it is even, then add to prefix
			if count%2 == 0 {
				palindrome += count * 2
			} else {
				center = 2
				palindrome += (count - 1) * 2
			}

			continue
		}

		// check if word has counter part
		reversed := reverse(word)
		if reversedCount, exists := freq[reversed]; exists {
			palindrome += min(count, reversedCount) * 2
		}
	}

	return palindrome + center
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// since all words are 2 letter words
func reverse(word string) string {
	bytes := []byte(word)
	bytes[0], bytes[1] = bytes[1], bytes[0]

	return string(bytes)
}

// since all words are 2 letter words
func isPalindrome(word string) bool {
	return word[0] == word[1]
}
