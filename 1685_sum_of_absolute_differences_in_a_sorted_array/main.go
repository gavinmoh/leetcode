package main

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)

	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	output := make([]int, n)

	for i, num := range nums {
		rightSum := totalSum - leftSum - num

		leftCount := i
		rightCount := n - 1 - i

		leftTotal := leftCount*num - leftSum
		rightTotal := rightSum - rightCount*num
		output[i] = leftTotal + rightTotal
		leftSum += num
	}

	return output

}
