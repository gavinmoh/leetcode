package main

func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	prevRow := []string{}
	for col := 0; col < n+1; col++ {
		prevRow = append(prevRow, str2[0:col])
	}

	for row := 1; row < m+1; row++ {
		currRow := make([]string, n+1)
		currRow[0] = str1[0:row]

		for col := 1; col <= n; col++ {
			if str1[row-1] == str2[col-1] {
				currRow[col] = string(prevRow[col-1]) + string(str1[row-1])
			} else {
				pickS1 := prevRow[col]
				pickS2 := currRow[col-1]
				if len(pickS1) < len(pickS2) {
					currRow[col] = pickS1 + string(str1[row-1])
				} else {
					currRow[col] = pickS2 + string(str2[col-1])
				}
			}
		}
		prevRow = currRow
	}

	return prevRow[n]
}
