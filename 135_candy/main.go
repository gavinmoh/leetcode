package main

func Candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	} else if len(ratings) == 1 {
		return 1
	}

	// initialize candy count to 1
	// because each child must have at least one candy
	candies := make([]int, len(ratings))
	for i, _ := range candies {
		candies[i] = 1
	}

	// traverse from left to right
	for i := 1; i < len(ratings); i++ {
		// if rating is higher than previous child
		// give one more candy than previous child
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// traverse from right to left
	for i := len(ratings) - 2; i >= 0; i-- {
		// if current child rating is higher than next child
		// and current child has less or equal candy than next child
		// give one more candy to current child
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}

	candyCount := 0
	for _, count := range candies {
		candyCount += count
	}

	return candyCount
}
