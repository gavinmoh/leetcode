package main

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	l := len(n)
	result := []int{int(math.Pow10(l-1)) - 1, int(math.Pow10(l)) + 1}
	mid := (l + 1) / 2
	left, _ := strconv.Atoi(n[:mid])
	for _, num := range []int{left - 1, left, left + 1} {
		temp := num
		if l&1 == 1 {
			temp /= 10
		}
		for temp > 0 {
			num = num*10 + temp%10
			temp /= 10
		}
		result = append(result, num)
	}
	output := -1
	num, _ := strconv.Atoi(n)
	for _, palindrome := range result {
		if palindrome != num {
			if output == -1 || abs(palindrome-num) < abs(output-num) || abs(palindrome-num) == abs(output-num) && palindrome < output {
				output = palindrome
			}
		}
	}
	return strconv.Itoa(output)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
