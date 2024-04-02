package main

func isIsomorphic(s string, t string) bool {
	sMap := [256]rune{}
	tMap := [256]rune{}

	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}
		sMap[s[i]] = rune(i + 1)
		tMap[t[i]] = rune(i + 1)
	}

	return true
}
