package main

func firstUniqChar(s string) int {
	freqMap := make(map[rune][]int)

	for i, char := range s {
		if _, ok := freqMap[char]; !ok {
			arr := make([]int, 2)
			arr[0], arr[1] = i, 1
			freqMap[char] = arr
		} else {
			freqMap[char][1]++
		}
	}

	index := -1
	for _, freq := range freqMap {
		if freq[1] > 1 {
			continue
		}
		if index == -1 || index > freq[0] {
			index = freq[0]
		}
	}

	return index
}
