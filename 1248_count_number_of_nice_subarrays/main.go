package main

func numberOfSubarrays(nums []int, k int) int {
	oddCount := 0
	count := 0
	left, mid := 0, 0
	for _, num := range nums {
		if num%2 != 0 {
			oddCount++
		}
		for oddCount > k {
			if isOdd(nums[left]) {
				oddCount--
			}
			left++
			mid = left
		}

		if oddCount == k {
			for !isOdd(nums[mid]) {
				mid++
			}
			count += (mid - left) + 1
		}
	}

	return count
}

func isOdd(x int) bool {
	return x%2 != 0
}
