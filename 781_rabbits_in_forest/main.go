package main

func numRabbits(answers []int) int {
	freq := map[int]int{}
	for _, answer := range answers {
		freq[answer]++
	}

	result := 0
	for colour, count := range freq {
		groups := count / (colour + 1)
		if count%(colour+1) != 0 {
			groups += 1
		}

		result += groups * (colour + 1)
	}

	return result
}
