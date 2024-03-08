package main

func maxFrequencyElements(nums []int) int {
	freqMap := make(map[int]int)
	maxFreq := 0
	result := 0

	for _, num := range nums {
		freqMap[num]++
		count := freqMap[num]
		if freqMap[num] > maxFreq {
			maxFreq = freqMap[num]
			result = count
		} else if freqMap[num] == maxFreq {
			result += count
		}
	}

	return result
}
