package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	score := 0
	left, right := 0, len(tokens)-1

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			score++
			left++
		} else if score > 0 && right-left > 1 {
			power += tokens[right]
			score--
			right--
		} else {
			break
		}
	}

	return score
}
