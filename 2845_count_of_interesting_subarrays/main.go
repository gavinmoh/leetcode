package main

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	count := make(map[int]int)
	count[0] = 1
	result := int64(0)
	prefix := 0
	for _, num := range nums {
		if num%modulo == k {
			prefix++
		}

		result += int64(count[(prefix-k+modulo)%modulo])
		count[prefix%modulo]++
	}

	return result
}
