package main

func threeConsecutiveOdds(arr []int) bool {
	for left, num := range arr {
		if num%2 == 0 {
			continue
		}
		count := 0
		for right := left + 1; right < left+3 && right < len(arr); right++ {
			if arr[right]%2 == 0 {
				break
			}
			count++
			if count >= 2 {
				return true
			}
		}
	}

	return false
}
