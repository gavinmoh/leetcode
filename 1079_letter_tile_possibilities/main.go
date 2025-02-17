package main

func numTilePossibilities(tiles string) int {
	used := make([]bool, len(tiles))
	sequences := map[string]int{}

	var dfs func(str string)
	dfs = func(str string) {
		if len(str) != 0 {
			sequences[str]++
		}

		if len(str) == len(tiles) {
			return
		}

		for i, char := range tiles {
			if used[i] {
				continue
			}

			used[i] = true

			dfs(str + string(char))

			used[i] = false
		}
	}

	dfs("")

	return len(sequences)
}
