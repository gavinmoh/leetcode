package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func countGoodIntegers(n int, k int) int64 {
	palindromes := map[string]bool{}
	base := int(math.Pow10((n - 1) / 2))
	skip := n & 1

	// get all the k palindromes
	for i := base; i < base*10; i++ {
		s := strconv.Itoa(i)
		rev := reverse(s)
		s += rev[skip:]

		num, _ := strconv.ParseInt(s, 10, 64)
		if num%int64(k) == 0 {
			// rearrange and sort chars in S to ensure uniq combinations
			chars := strings.Split(s, "")
			sort.Strings(chars)
			palindromes[strings.Join(chars, "")] = true
		}
	}

	factorial := make([]int64, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * int64(i)
	}

	count := int64(0)
	for s := range palindromes {
		freq := make([]int, 10)
		for _, c := range s {
			freq[c-'0']++
		}

		// calculate permutations and combinations
		tot := int64(n-freq[0]) * factorial[n-1]
		for _, x := range freq {
			tot /= factorial[x]
		}
		count += tot
	}

	return count
}

func reverse(s string) string {
	runes := []rune(s)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	return string(runes)
}
