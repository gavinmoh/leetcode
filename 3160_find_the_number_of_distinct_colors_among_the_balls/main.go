package main

func queryResults(limit int, queries [][]int) []int {
	balls := map[int]int{} // using slice will lead to memory insufficient error
	colourFreq := map[int]int{}
	result := []int{}

	for _, query := range queries {
		i, colour := query[0], query[1]
		originalColour := 0
		if _, ok := balls[i]; ok {
			originalColour = balls[i]
		}

		if originalColour != 0 {
			if colourFreq[originalColour] == 1 {
				delete(colourFreq, originalColour)
			} else {
				colourFreq[originalColour]--
			}
		}

		// paint the ball
		balls[i] = colour
		colourFreq[colour]++

		result = append(result, len(colourFreq))
	}

	return result
}
