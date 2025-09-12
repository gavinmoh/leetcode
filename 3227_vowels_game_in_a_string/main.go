package main

func doesAliceWin(s string) bool {
	for _, c := range s {
		if isVowel(c) {
			return true
		}
	}

	return false
}

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}
