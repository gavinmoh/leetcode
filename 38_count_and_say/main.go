package main

import "strings"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	recur := countAndSay(n - 1)
	pairs := helper1(recur)

	return helper2(pairs)
}

func helper1(s string) [][]int {
	pairs := [][]int{}
	digit := s[0]
	count := 1

	for i := 1; i < len(s); i++ {
		if digit == s[i] {
			count++
		} else {
			pairs = append(pairs, []int{count, int(digit - '0')})
			digit = s[i]
			count = 1
		}
	}
	pairs = append(pairs, []int{count, int(digit - '0')})

	return pairs
}

func helper2(pairs [][]int) string {
	var result strings.Builder
	for _, pair := range pairs {
		count, digit := pair[0], pair[1]
		result.WriteRune(rune('0' + count))
		result.WriteRune(rune('0' + digit))
	}

	return result.String()
}
