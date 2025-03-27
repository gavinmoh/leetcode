package main

func minimumIndex(nums []int) int {
	n := len(nums)
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	x := -1
	for num, count := range freq {
		if count > freq[x] {
			x = num
		}
	}

	leftCount, f := 0, freq[x]
	for i, num := range nums {
		if num == x {
			leftCount++

			rightCount := f - leftCount
			leftLength := i - 0 + 1
			rightLength := n - i - 1

			if leftCount > leftLength/2 && rightCount > rightLength/2 {
				return i
			}
		}
	}

	return -1
}
