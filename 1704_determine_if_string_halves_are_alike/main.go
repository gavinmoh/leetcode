package main

func halvesAreAlike(s string) bool {
	half := len(s) / 2
	a := s[:half]
	b := s[half:]

	aCount := 0
	bCount := 0

	for i := 0; i < half; i++ {
		if isVowel(a[i]) {
			aCount += 1
		}
		if isVowel(b[i]) {
			bCount += 1
		}
	}

	return aCount == bCount
}

func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
