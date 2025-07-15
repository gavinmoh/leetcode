package main

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	vowelCount, consonantCount := 0, 0
	for _, c := range word {
		if isDigit(c) {
			continue
		}

		if isLetter(c) {
			if isVowel(c) {
				vowelCount++
			} else {
				consonantCount++
			}

			continue
		}

		return false
	}

	if vowelCount <= 0 {
		return false
	}

	if consonantCount <= 0 {
		return false
	}

	return true
}

func isVowel(c rune) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}

func isLetter(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= 'A' && c <= 'Z' {
		return true
	}

	return false
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	return false
}
