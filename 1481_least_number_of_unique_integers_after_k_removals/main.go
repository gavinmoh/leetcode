package main

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	freqMap := make(map[int]int)
	uniqueNums := []int{}

	for _, num := range arr {
		if _, ok := freqMap[num]; !ok {
			uniqueNums = append(uniqueNums, num)
		}
		freqMap[num]++
	}

	sort.SliceStable(uniqueNums, func(i, j int) bool {
		return freqMap[uniqueNums[i]] < freqMap[uniqueNums[j]]
	})

	uniqueCount := len(uniqueNums)
	for _, num := range uniqueNums {
		if k >= freqMap[num] {
			k -= freqMap[num]
			uniqueCount--
		} else {
			break
		}
	}

	return uniqueCount
}
