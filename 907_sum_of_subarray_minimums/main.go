package main

func sumSubarrayMins(arr []int) int {
	modulo := 1_000_000_007
	sum := 0
	stack := [][]int{}

	for i, num := range arr {
		for len(stack) > 0 && num < stack[len(stack)-1][1] {
			pair := stack[len(stack)-1]
			j, val := pair[0], pair[1]
			left, right := j+1, i-j
			stack = stack[:len(stack)-1]

			if len(stack) > 0 {
				left = j - stack[len(stack)-1][0]
			}

			sum = (sum + val*left*right) % modulo
		}
		stack = append(stack, []int{i, num})
	}

	for i := 0; i < len(stack); i++ {
		j, val := stack[i][0], stack[i][1]
		left, right := j+1, len(arr)-j

		if i > 0 {
			left = j - stack[i-1][0]
		}
		sum = (sum + val*left*right) % modulo
	}

	return sum
}
