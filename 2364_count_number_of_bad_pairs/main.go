package main

func countBadPairs(nums []int) int64 {
	n := int64(len(nums))
	totalPairs := n * (n - 1) / 2
	counts := map[int]int{}

	for i := 0; i < len(nums); i++ {
		counts[nums[i]-i]++
	}

	goodPairs := int64(0)
	for _, count := range counts {
		if count < 2 {
			continue
		}
		goodPairs += int64(count * (count - 1) / 2)
	}

	return totalPairs - goodPairs
}
