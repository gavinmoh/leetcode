package main

func minChanges(s string) int {
	result := 0

	if len(s) > 2 {
		// break down s into block of 2
		for i := 0; i < len(s); i += 2 {
			result += minChanges(s[i : i+2])
		}
	} else if s == "10" || s == "01" {
		result += 1
	}

	return result
}
