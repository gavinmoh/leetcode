package main

func maxBottlesDrunk(numBottles int, numExchange int) int {
	full, empty := 0, numBottles
	drunk := numBottles
	for full > 0 || empty >= numExchange {
		// drink all full
		drunk += full
		empty += full
		full = 0

		// exchange
		for empty >= numExchange {
			full += 1
			empty -= numExchange
			numExchange++
		}
	}

	return drunk
}
