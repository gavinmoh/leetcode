package main

func getWinner(arr []int, k int) int {
	if k > len(arr) {
		max := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}

	winner := arr[0]
	winCount := 0

	for _, i := range arr[1:] {
		if i > winner {
			winner = i
			winCount = 1
		} else {
			winCount++
		}

		if winCount == k {
			break
		}
	}

	return winner
}
