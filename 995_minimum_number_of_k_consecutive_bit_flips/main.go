package main

func minKBitFlips(nums []int, k int) int {
	queue := []int{}
	count := 0

	for i, num := range nums {
		for len(queue) > 0 && i > queue[0]+k-1 {
			queue = queue[1:]
		}

		if (num+len(queue))%2 == 0 {
			if i+k > len(nums) {
				return -1
			}
			count++
			queue = append(queue, i)
		}
	}

	return count
}
