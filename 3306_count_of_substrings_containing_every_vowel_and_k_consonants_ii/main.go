package main

func countOfSubstrings(word string, k int) int64 {
	vowelCount := map[byte]int{}
	consonantCount := 0

	nextConsonant, nextConsonantIdx := make([]int, len(word)), len(word)
	for i := len(word) - 1; i >= 0; i-- {
		nextConsonant[i] = nextConsonantIdx
		if !isVowel(word[i]) {
			nextConsonantIdx = i
		}
	}

	left, right := 0, 0
	result := int64(0)
	for right < len(word) {
		c := word[right]
		if isVowel(c) {
			vowelCount[c]++
		} else {
			consonantCount++
		}

		for consonantCount > k {
			leftChar := word[left]
			if isVowel(leftChar) {
				vowelCount[leftChar]--
				if vowelCount[leftChar] == 0 {
					delete(vowelCount, leftChar)
				}
			} else {
				consonantCount--
			}

			left++
		}

		for left < len(word) && len(vowelCount) == 5 && consonantCount == k {
			result += int64(nextConsonant[right] - right)
			leftChar := word[left]
			if isVowel(leftChar) {
				vowelCount[leftChar]--
				if vowelCount[leftChar] == 0 {
					delete(vowelCount, leftChar)
				}
			} else {
				consonantCount--
			}

			left++
		}

		right++
	}

	return result
}

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}
