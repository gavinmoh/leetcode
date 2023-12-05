package main

func numberOfMatches(n int) int {
	if n == 2 {
		return 1
	}

	// this code will result in TLE error
	// if n%2 == 0 {
	// 	return (n / 2) + numberOfMatches(n/2)
	// }

	// currentMatches := (n - 1) / 2
	// teamsForNextRound := currentMatches + 1
	// return currentMatches + numberOfMatches(teamsForNextRound)

	return n - 1
}
