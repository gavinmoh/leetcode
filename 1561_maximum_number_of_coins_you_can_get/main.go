package main

import "sort"

func maxCoins(piles []int) int {

	sort.Ints(piles)

	coins := 0

	// remove the piles for Bob, which will always be the smallest
	piles = piles[len(piles)/3:]

	// always take the smaller pile, i.e. the odd piles
	for i := 0; i < len(piles); i += 2 {
		coins += piles[i]
	}

	return coins
}
