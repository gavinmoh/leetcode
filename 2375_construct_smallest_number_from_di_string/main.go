package main

import "strings"

func smallestNumber(pattern string) string {
	n := len(pattern)
	num := make([]int, n+1)
	used := make([]bool, n+2)

	minNum := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		minNum[i] = 10
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == n+1 {
			if smaller(num, minNum) {
				copy(minNum, num)
			}
			return
		}

		p := pattern[i-1]

		for digit := 1; digit <= n+1; digit++ {
			if used[digit] {
				continue
			}

			if p == 'I' && digit < num[i-1] {
				continue
			}

			if p == 'D' && digit > num[i-1] {
				continue
			}

			used[digit] = true

			num[i] = digit
			dfs(i + 1)

			used[digit] = false
		}
	}

	for digit := 1; digit <= n+1; digit++ {
		num[0] = digit
		used[digit] = true
		dfs(1)
		used[digit] = false
		num[0] = 10
	}

	var result strings.Builder
	for _, digit := range minNum {
		result.WriteRune('0' + rune(digit))
	}

	return result.String()
}

func smaller(arr1, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			return false
		}
	}

	return true
}
