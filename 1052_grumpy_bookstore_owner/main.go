package main

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	currentSatisfied := 0
	for i := 0; i < len(customers) && i < minutes; i++ {
		currentSatisfied += customers[i]
	}
	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 0 {
			currentSatisfied += customers[i]
		}
	}

	maxSatisfied := currentSatisfied
	for i := minutes; i < len(customers); i++ {
		if grumpy[i-minutes] == 1 {
			currentSatisfied -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			currentSatisfied += customers[i]
		}
		if currentSatisfied > maxSatisfied {
			maxSatisfied = currentSatisfied
		}
	}

	return maxSatisfied
}
