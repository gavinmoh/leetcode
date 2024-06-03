package main

func appendCharacters(s string, t string) int {
	m, n := len(s), len(t)
	left, right := 0, 0

	for left < m && right < n {
		if s[left] == t[right] {
			right++
		}
		left++
	}

	return len(t) - right
}
