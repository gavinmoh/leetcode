package main

func rearrangeArray(nums []int) []int {
	positives := []int{}
	negatives := []int{}

	for _, num := range nums {
		if num > 0 {
			positives = append(positives, num)
		} else {
			negatives = append(negatives, num)
		}
	}

	pos, neg := 0, 0
	for i := 0; i < len(nums); i++ {
		if (i % 2) == 0 {
			nums[i] = positives[pos]
			pos++
		} else {
			nums[i] = negatives[neg]
			neg++
		}
	}

	return nums
}
