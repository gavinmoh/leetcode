package main

func LastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}

	var sortArray = func(arr []int) []int {
		for i := 0; i < len(arr); i++ {
			swapped := false
			for j := i; j < len(arr); j++ {
				if arr[i] <= arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
					swapped = true
				}
			}
			if !swapped {
				break
			}
		}
		return arr
	}

	sorted := sortArray(stones)
	if len(sorted) == 2 {
		return sorted[0] - sorted[1]
	}
	if sorted[0] == sorted[1] {
		return LastStoneWeight(sorted[2:])
	}
	newWeight := sorted[0] - sorted[1]
	sorted = append(sorted[2:], newWeight)
	return LastStoneWeight(sorted)
}
