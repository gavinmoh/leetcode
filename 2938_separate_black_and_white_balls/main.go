package main

func minimumSteps(s string) int64 {
	result := int64(0)
	count := int64(0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			count++
		} else {
			result += count
		}
	}

	return result
}
