package main

func minLength(s string) int {
	for {
		occured := false
		for i := 1; i < len(s); i++ {
			if s[i-1:i+1] == "AB" || s[i-1:i+1] == "CD" {
				s = s[:i-1] + s[i+1:]
				occured = true
				break
			}
		}
		if !occured {
			break
		}
	}

	return len(s)
}
