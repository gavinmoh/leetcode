package main

func kLengthApart(nums []int, k int) bool {
	last1 := -k - 1
	for i, num := range nums {
		if num == 0 {
			continue
		}

		if i-last1-1 < k {
			return false
		}
		last1 = i
	}

	return true
}
