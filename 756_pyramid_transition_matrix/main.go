package main

func pyramidTransition(bottom string, allowed []string) bool {
	allowedMap := map[string][]byte{}
	for _, block := range allowed {
		prefix := block[0:2]
		allowedMap[prefix] = append(allowedMap[prefix], block[2])
	}

	var solve func(string) bool
	solve = func(b string) bool {
		if len(b) == 1 {
			return true
		}

		possibles := make([][]byte, len(b)-1)
		for i := range possibles {
			possibles[i] = []byte{}
		}

		for i := 0; i < len(b)-1; i++ {
			prefix := b[i : i+2]
			if _, ok := allowedMap[prefix]; !ok {
				return false
			}

			possibles[i] = allowedMap[prefix]
		}

		next := []string{}
		var dfs func(i int, curr string)
		dfs = func(i int, curr string) {
			if i >= len(possibles) {
				next = append(next, curr)
				return
			}

			for _, c := range possibles[i] {
				dfs(i+1, curr+string(c))
			}
		}

		dfs(0, "")

		for _, n := range next {
			if solve(n) {
				return true
			}
		}

		return false
	}

	return solve(bottom)
}
