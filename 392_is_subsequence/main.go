package main

func isSubsequence(s string, t string) bool {
	i := 0
	for _, c := range t {
		if i < len(s) && c == rune(s[i]) { // check if c is equal to the current character in s
			i++ // move to the next character in s
		}
	}
	return i == len(s) // if i is equal to the length of s, then s is a subsequence of t
}
