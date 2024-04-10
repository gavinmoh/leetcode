package main

import "sort"

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	var dfs func([]int) []int
	dfs = func(cards []int) []int {
		n := len(cards)
		if n == 1 {
			return cards
		}

		// split half
		half := n / 2
		var top []int
		var bottom []int
		if n%2 == 0 {
			top = cards[:half]
			bottom = dfs(cards[half:])
		} else {
			half += 1
			top = cards[:half]
			bottom = dfs(cards[half:])
			bottom = append(bottom[len(bottom)-1:], bottom[:len(bottom)-1]...)
		}

		reordered := []int{}
		for i := 0; i < half; i++ {
			reordered = append(reordered, top[i])
			if i < len(bottom) {
				reordered = append(reordered, bottom[i])
			}
		}
		return reordered
	}

	return dfs(deck)
}
