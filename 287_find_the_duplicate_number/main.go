package main

func findDuplicate(nums []int) int {

	// floyd's cycle detection algorithm
	// find the intersection point of the fast and slow pointers
	slowIndex := 0
	fastIndex := 0
	for {
		slowIndex = nums[slowIndex]
		fastIndex = nums[nums[fastIndex]]

		if slowIndex == fastIndex {
			break
		}
	}

	// find the "entrance" to the cycle
	finderIndex := 0
	for {
		finderIndex = nums[finderIndex]
		slowIndex = nums[slowIndex]

		if finderIndex == slowIndex {
			break
		}
	}

	return finderIndex

}
