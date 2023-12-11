package main

func findSpecialInteger(arr []int) int {
	freqMap := make(map[int]int, len(arr))
	maxFreq := 0
	maxNum := arr[0]

	for _, num := range arr {
		freqMap[num]++
		if freqMap[num] > maxFreq {
			maxFreq = freqMap[num]
			maxNum = num
		}
	}

	return maxNum
}
