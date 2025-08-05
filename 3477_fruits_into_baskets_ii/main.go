package main

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	count := 0
	used := make([]bool, n)

	for _, quantity := range fruits {
		for basket, capacity := range baskets {
			if used[basket] || quantity > capacity {
				continue
			}

			used[basket] = true
			count++
			break
		}
	}

	return n - count
}
