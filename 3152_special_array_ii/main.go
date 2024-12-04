package main

func isArraySpecial(nums []int, queries [][]int) []bool {
	prefix := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		if nums[i]%2 == nums[i-1]%2 {
			prefix[i] = prefix[i-1] + 1
		} else {
			prefix[i] = prefix[i-1]
		}
	}

	answers := []bool{}
	for _, query := range queries {
		answers = append(answers, prefix[query[1]]-prefix[query[0]] == 0)
	}

	return answers
}
