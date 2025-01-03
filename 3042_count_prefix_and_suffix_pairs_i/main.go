package main

func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}

	return count
}

func isPrefixAndSuffix(str1, str2 string) bool {
	if len(str1) > len(str2) {
		return false
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}

		if str1[i] != str2[len(str2)-len(str1)+i] {
			return false
		}
	}

	return true
}
