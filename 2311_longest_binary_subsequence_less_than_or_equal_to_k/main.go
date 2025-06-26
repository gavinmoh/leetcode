package main

import "strconv"

func longestSubsequence(s string, k int) int {
	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			binaryStr := s[i:]
			num, _ := strconv.ParseInt(binaryStr, 2, 64)
			if num <= int64(k) {
				result++
			}
		} else {
			result++
		}
	}

	return result
}
