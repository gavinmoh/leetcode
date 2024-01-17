package main

import "sort"

func uniqueOccurrences(arr []int) bool {
	freqMap := make(map[int]int)

	for _, num := range arr {
		freqMap[num]++
	}

	frequencies := []int{}
	for _, freq := range freqMap {
		frequencies = append(frequencies, freq)
	}

	sort.Ints(frequencies)
	for i := 0; i < len(frequencies)-1; i++ {
		if frequencies[i] == frequencies[i+1] {
			return false
		}
	}
	return true
}
