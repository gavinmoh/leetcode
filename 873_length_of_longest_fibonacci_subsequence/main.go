package main

func lenLongestFibSubseq(arr []int) int {
	nums := map[int]bool{}
	for _, num := range arr {
		nums[num] = true
	}

	maxLength := 0
	for left := 0; left < len(arr)-1; left++ {
		for right := left + 1; right < len(arr); right++ {
			prev := arr[right]
			curr := arr[left] + arr[right]
			currLength := 2

			for {
				if _, ok := nums[curr]; !ok {
					break
				}

				prev, curr = curr, curr+prev
				currLength++
				maxLength = max(maxLength, currLength)
			}
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
