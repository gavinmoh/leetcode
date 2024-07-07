package main

func numWaterBottles(numBottles int, numExchange int) int {
	emptyBottles := 0
	fullBottles := numBottles
	count := 0

	for emptyBottles+fullBottles >= numExchange {
		// drink water
		count += fullBottles
		emptyBottles += fullBottles
		fullBottles = 0

		// exchange
		fullBottles = emptyBottles / numExchange
		emptyBottles %= numExchange
	}
	count += fullBottles

	return count
}
