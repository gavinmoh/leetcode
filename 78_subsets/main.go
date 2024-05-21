package main

func subsets(nums []int) [][]int {
	result := [][]int{{}}

	for _, num := range nums {
		newSubsets := [][]int{}
		for _, subset := range result {
			temp := append([]int{}, subset...)
			temp = append(temp, num)
			newSubsets = append(newSubsets, temp)
		}
		for _, subset := range newSubsets {
			result = append(result, subset)
		}
	}

	return result
}
