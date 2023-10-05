package main

func majorityElement(nums []int) []int {
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
	}

	oneThird := float32(len(nums)) / 3.0

	var result []int
	for num, frequency := range frequencyMap {
		if float32(frequency) > oneThird {
			result = append(result, num)
		}
	}

	return result

}
