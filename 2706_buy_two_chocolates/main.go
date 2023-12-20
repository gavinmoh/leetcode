package main

import "sort"

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	leftOver := money - (prices[0] + prices[1])

	if leftOver < 0 {
		return money
	}

	return leftOver
}
