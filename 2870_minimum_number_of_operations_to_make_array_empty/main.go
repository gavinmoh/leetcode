package main

func minOperations(nums []int) int {
	freqMap := make(map[int]int)
	cache := make(map[int]int)

	for _, num := range nums {
		if _, ok := freqMap[num]; !ok {
			freqMap[num] = 1
		} else {
			freqMap[num]++
		}
	}

	count := 0
	for _, freq := range freqMap {
		if cachedResult, ok := cache[freq]; ok {
			count += cachedResult
			continue
		}

		result := helper(freq, 0)
		if result == -1 {
			return -1
		}
		cache[freq] = result
		count += result
	}

	return count
}

func helper(num int, acc int) int {
	switch num {
	case 0:
		return 0 + acc
	case 1:
		return -1
	case 2:
		return 1 + acc
	case 3:
		return 1 + acc
	case 4:
		return 2 + acc
	default:
		return helper(num-3, acc+1)
	}
}
