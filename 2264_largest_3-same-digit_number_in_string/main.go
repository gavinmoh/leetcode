package main

func largestGoodInteger(num string) string {
	output := ""

	for i := 2; i < len(num); i++ {
		if !(num[i-2] == num[i-1] && num[i-1] == num[i]) {
			continue
		}

		substr := num[i-2 : i+1]
		if output == "" {
			output = substr
			continue
		}

		output = max(output, substr)
	}

	return output
}

func max(a, b string) string {
	if a > b {
		return a
	}
	return b
}
