package main

func answerString(word string, numFriends int) string {
	// edge cases, where the word cannot be split at all
	if numFriends == 1 {
		return word
	}

	n := len(word)
	result := ""
	for left := 0; left < n; left++ {
		right := left + n - numFriends + 1
		result = larger(result, word[left:min(right, n)])
	}

	return result
}

func larger(a, b string) string {
	for i := 0; i < (min(len(a), len(b))); i++ {
		if a[i] > b[i] {
			return a
		} else if b[i] > a[i] {
			return b
		}
	}

	if len(a) > len(b) {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
