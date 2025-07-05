package main

func findLucky(arr []int) int {
	freq := map[int]int{}
	for _, num := range arr {
		freq[num]++
	}

	largest := -1
	for num, count := range freq {
		if num == count && num > largest {
			largest = num
		}
	}

	return largest
}
