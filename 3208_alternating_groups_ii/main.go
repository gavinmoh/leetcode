package main

func numberOfAlternatingGroups(colors []int, k int) int {
	for i := 0; i < k-1; i++ {
		colors = append(colors, colors[i])
	}

	n := len(colors)
	result := 0
	left, right := 0, 1

	for right < n {
		// check if prev tile same as current
		if colors[right] == colors[right-1] {
			// reset window
			left = right
			right++
			continue
		}

		right++

		if (right - left) < k {
			continue
		}

		result++
		left++ // shrink window
	}

	return result
}
