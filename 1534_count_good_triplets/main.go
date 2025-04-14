package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	count := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}

			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) > b {
					continue
				}

				if abs(arr[i]-arr[k]) > c {
					continue
				}

				count++
			}
		}
	}

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
