package main

func countSubstrings(s string) int {
	n := len(s)
	result := 0

	helper := func(left, right int) int {
		count := 0
		for left >= 0 && right < n && s[left] == s[right] {
			count += 1
			left -= 1
			right += 1
		}
		return count
	}

	for i := 0; i < n; i++ {
		result += helper(i, i)
		result += helper(i, i+1)
	}

	return result
}
