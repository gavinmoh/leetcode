package main

func findErrorNums(nums []int) []int {
	missing, repeated := 0, 0
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	for i := 1; i <= len(nums); i++ {
		freq, ok := freqMap[i]
		if ok && freq > 1 {
			repeated = i
		} else if !ok {
			missing = i
		}
		if missing != 0 && repeated != 0 {
			break
		}
	}

	return []int{repeated, missing}
}
