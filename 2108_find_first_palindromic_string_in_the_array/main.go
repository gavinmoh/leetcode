package main

func firstPalindrome(words []string) string {
	for _, word := range words {
		isPalindromic := true
		left, right := 0, len(word)-1
		for right > left {
			if word[left] != word[right] {
				isPalindromic = false
				break
			}
			left++
			right--
		}

		if isPalindromic {
			return word
		}
	}

	return ""
}
