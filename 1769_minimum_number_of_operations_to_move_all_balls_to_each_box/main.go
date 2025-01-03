package main

func minOperations(boxes string) []int {
	result := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		// balls from left
		for j := 0; j < (i - 0); j++ {
			if boxes[j] == '1' {
				result[i] += abs(i - j)
			}
		}

		// balls from right
		for j := i + 1; j < len(boxes); j++ {
			if boxes[j] == '1' {
				result[i] += abs(i - j)
			}
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
