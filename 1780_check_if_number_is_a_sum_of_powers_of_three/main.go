package main

import "strconv"

func checkPowersOfThree(n int) bool {
	base3Str := strconv.FormatInt(int64(n), 3)
	for _, char := range base3Str {
		if char == '2' {
			return false
		}
	}

	return true
}
