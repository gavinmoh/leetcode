package main

func checkRecord(s string) bool {
	n := len(s)

	var dfs func(i int, absentCount int, consecutiveLateCount int) bool
	dfs = func(i int, absentCount int, consecutiveLateCount int) bool {
		if absentCount >= 2 || consecutiveLateCount >= 3 {
			return false
		}

		if i >= n {
			return true
		}

		switch s[i] {
		case 'A':
			return dfs(i+1, absentCount+1, 0)
		case 'L':
			return dfs(i+1, absentCount, consecutiveLateCount+1)
		default:
			return dfs(i+1, absentCount, 0)
		}
	}

	return dfs(0, 0, 0)
}
