package main

func maximumSum(nums []int) int {
	digitSums := map[int][]int{}

	for _, num := range nums {
		sum := digitSum(num)
		if _, ok := digitSums[sum]; !ok {
			digitSums[sum] = []int{}
		}
		digitSums[sum] = append(digitSums[sum], num)
	}

	maxSum := -1
	for _, nums := range digitSums {
		if len(nums) < 2 {
			continue
		}

		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				sum := nums[i] + nums[j]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}

	return maxSum
}

func digitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += (x % 10)
		x /= 10
	}

	return sum
}
