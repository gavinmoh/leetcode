package main

func buildArray(target []int, n int) []string {
	stack := []string{}
	lastNum := target[len(target)-1]
	index := 0

	for i := 1; i <= n; i++ {
		// since target is strictly increasing,
		// we can just check if i is greater than the last number in target
		if i > lastNum {
			break
		}

		if i == target[index] {
			stack = append(stack, "Push")
			index++
		} else {
			stack = append(stack, "Push", "Pop")
		}
	}

	return stack
}
