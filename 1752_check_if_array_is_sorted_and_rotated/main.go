package main

func check(nums []int) bool {
	n, x := len(nums), 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			x = i
			break
		}
	}

	nums = append(nums[x:n], nums[:x]...)
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}

	return true
}
