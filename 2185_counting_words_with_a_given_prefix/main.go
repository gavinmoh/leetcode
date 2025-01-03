package main

func prefixCount(words []string, pref string) int {
	count := 0
	for _, word := range words {
		if isPrefix(pref, word) {
			count++
		}
	}

	return count
}

func isPrefix(str1, str2 string) bool {
	if len(str1) > len(str2) {
		return false
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}

	return true
}
