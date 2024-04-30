package main

func wonderfulSubstrings(word string) int64 {
	n := len(word)
	state := 0
	count := [1 << 10]int64{}
	count[state] += 1

	var result int64
	for i := 0; i < n; i++ {
		k := word[i] - 'a'
		state = state ^ (1 << k)

		result += count[state]

		for j := 0; j < 10; j++ {
			stateJ := state ^ (1 << j)
			result += count[stateJ]
		}

		count[state] += 1
	}

	return result
}
